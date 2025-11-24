`epub-tool` removes from epub files CSS styles that harm the reading experience.

Currently this means:
  1. Removing font size declarations that make it impossible to adjust the font size in your ebook reader
  2. Removing text color declarations that cause the text color to remain black when the reader switches the background to black in dark mode
  3. Optionally (using the --removeBackgroundColors flag) removing background color declarations that might make some graphic elements invisible in dark mode

To run the basic cleanup: epub-tool <target-epub-file>

To additionally remove all background colors: epub-tool -b <target-epub-file>

The command outputs a new file which is a copy of <target-epub-file>, but with the chosen styles removed and the prefix "_cleaned" added to its filename. The output filename can be customised with the -o flag.

Usage:
  `epub-tool [flags]`

Flags:

  `-d`, `--dryRun`                  print changes that would be made, but don't write them to disk
  
  `-h`, `--help`                    help for epub-tool
  
  `-o`, `--outputFileName string`   name of the cleaned output file
  
  `-v`, `--verbose`                 enable verbose output