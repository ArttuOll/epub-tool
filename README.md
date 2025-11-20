`epub-tool` removes from epub files CSS styles that harm the reading experience.

Currently this means:
  1. Removing font size declarations that make it impossible to adjust the font size in your ebook reader
  2. Removing text color declarations that cause the text color to remain black when the reader switches the background to black in dark mode

Usage:
  `epub-tool [flags]`

Flags:

  `-d`, `--dryRun`                  print changes that would be made, but don't write them to disk
  
  `-h`, `--help`                    help for epub-tool
  
  `-o`, `--outputFileName string`   name of the cleaned output file
  
  `-v`, `--verbose`                 enable verbose output