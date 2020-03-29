# Evgen

Google Calendar like event generator for Go.

## Features
- Daily, weekly, monthly, and yearly events
- Create events with number of occurrences or end date
- Weekly events with an option to choose multiple days in a week

## Usage
Generating daily events
```
startAt := time.Now()
frequency := 2 // every two days
qty := 4 // total of four occurrences

events := evegen.Generate("daily", startAt, nil, frequency, qty, nil)
```

Generating weekly events
```
startAt := time.Now()
dayOfWeek := []int{0, 2} // Sunday and Tuesday
frequency := 2 // every other week
qty := 0
endDate := startAt.AddDate(0, 4, 0) // four months from now

events := evegen.Generate("weekly", startAt, dayOfWeek, frequency, qty, endDate)
```

## TODO
 - Montly events with an option to choose a day and a week of the month 

## Contribution
Contributions welcome via Github pull requests and issues.

## License
This project is licensed under the MIT License. Please refer the License.txt file.
