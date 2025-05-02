package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/rickar/cal/v2"
)

func main() {
	// Parse command-line arguments
	startDateStr := flag.String("start", "", "Start date (YYYY-MM-DD)")
	endDateStr := flag.String("end", "", "End date (YYYY-MM-DD)")
	holidaysFile := flag.String("holidays", "holidays.txt", "Path to holidays file")
	flag.Parse()

	// Validate input
	if *startDateStr == "" || *endDateStr == "" {
		fmt.Println("Both start and end dates are required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", *startDateStr)
	if err != nil {
		fmt.Printf("Invalid start date: %v\n", err)
		os.Exit(1)
	}

	endDate, err := time.Parse("2006-01-02", *endDateStr)
	if err != nil {
		fmt.Printf("Invalid end date: %v\n", err)
		os.Exit(1)
	}

	// Make sure start date is before end date
	if startDate.After(endDate) {
		fmt.Println("Start date must be before end date")
		os.Exit(1)
	}

	// Read holidays from file
	holidayDates, err := readHolidayDates(*holidaysFile)
	if err != nil {
		fmt.Printf("Error reading holidays: %v\n", err)
		os.Exit(1)
	}

	// Create calendar
	c := cal.NewBusinessCalendar()

	// Calculate workdays
	workdays := getWorkdays(startDate, endDate, c, holidayDates)

	// Display results using lipgloss for formatting
	displayWorkdays(workdays, startDate, endDate)
}

func readHolidayDates(filePath string) (map[string]bool, error) {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("holidays file '%s' not found", filePath)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	holidays := make(map[string]bool)
	fileIsEmpty := true

	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) > 0 {
			fileIsEmpty = false
		}
		
		dateStrs := strings.Split(line, ",")
		
		for _, dateStr := range dateStrs {
			dateStr = strings.TrimSpace(dateStr)
			if dateStr == "" {
				continue
			}
			
			// Validate the date format
			_, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				return nil, fmt.Errorf("invalid date format '%s': %v", dateStr, err)
			}
			
			// Add to holiday map
			holidays[dateStr] = true
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Check if the file was empty or had no valid dates
	if fileIsEmpty || len(holidays) == 0 {
		return nil, fmt.Errorf("holidays file '%s' is empty or contains no valid dates", filePath)
	}

	return holidays, nil
}

func getWorkdays(start, end time.Time, c *cal.BusinessCalendar, holidays map[string]bool) []time.Time {
	var workdays []time.Time
	
	// Normalize dates to start at the beginning of the day
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())
	
	// Iterate through each day
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		// Format the date to check against holidays map
		dateStr := d.Format("2006-01-02")
		
		// Check if it's a workday and not a holiday
		if c.IsWorkday(d) && !holidays[dateStr] {
			workdays = append(workdays, d)
		}
	}
	
	return workdays
}

func displayWorkdays(workdays []time.Time, startDate, endDate time.Time) {
	// Style definitions using lipgloss
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1)

	dateStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#04B575"))

	countStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFAE5D")).
		Bold(true)

	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#5D8BF4")).
		Italic(true)

	// Print the title
	fmt.Println(titleStyle.Render(" Workdays "))
	fmt.Println()

	// Print date range
	fmt.Printf("%s %s to %s\n\n", 
		headerStyle.Render("Date Range:"),
		dateStyle.Render(startDate.Format("Jan 2, 2006")),
		dateStyle.Render(endDate.Format("Jan 2, 2006")))

	// Print the workdays
	for i, day := range workdays {
		fmt.Printf("%s %s\n", 
			countStyle.Render(fmt.Sprintf("%3d.", i+1)),
			dateStyle.Render(day.Format("Monday, January 2, 2006")))
	}

	// Print summary
	fmt.Println()
	fmt.Printf("%s %s\n", 
		headerStyle.Render("Total workdays:"),
		countStyle.Render(fmt.Sprintf("%d", len(workdays))))
}