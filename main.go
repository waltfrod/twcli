package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/ktr0731/go-fuzzyfinder"
)

var col1Style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#4f46e5"))
var col2Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#0ea5e9"))
var col3Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#e5e7eb"))
var linkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#14b8a6")).Underline(true)

var headerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#058030")).
	Bold(true).
	Align(lipgloss.Center)

var defStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#0285c7"))

func main() {
	browser := flag.Bool("browser", false, "Abrir documentación en el navegador web")
	flag.Parse()
	idx, err := fuzzyfinder.Find(Items, func(i int) string { return Items[i].String() })
	if err != nil {
		log.Fatal(err)
	}

	if *browser {
		fmt.Println(
			headerStyle.Render(
				fmt.Sprintf(
					"🔗 %s  %s  %s",
					Items[idx].Section,
					Items[idx].Subsection,
					Items[idx].URL,
				),
			),
		)
		if err := open(Items[idx].URL); err != nil {
			log.Fatal(err)
		}
	} else {
		printItem(Items[idx])
	}
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func maxw(h []string, b [][]string) int {
	m := 0
	for _, v := range h {
		m = max(m, utf8.RuneCountInString(v))
	}
	for _, v := range b {
		for _, vv := range v {
			m = max(m, utf8.RuneCountInString(vv))
		}
	}

	return m
}

func printItem(item Register) {
	fmt.Println("┃", defStyle.Render(fmt.Sprintf(" %s  %s", item.Section, item.Subsection)))
	fmt.Println("┃", linkStyle.Render("󰌷 "+item.URL))
	fmt.Println("┃", col3Style.Render("󰭷 "+item.Description))

	m := maxw(item.Header, item.Body) + 2

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderRow(true).
		BorderColumn(true).
		Width((m * len(item.Header)) + len(item.Header) + 1).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return headerStyle.Copy().Width(m)
			}
			switch col {
			case 0:
				return col1Style.Copy().Width(m)
			case 1:
				return col2Style.Copy().Width(m)
			default:
				return col3Style.Copy().Width(m)
			}
		}).
		Headers(item.Header...).
		Rows(item.Body...)

	fmt.Println(t)
}
