# go-pacbiohifi-visualization

- bins the sequenced pacbiohifi reads for visualization
- this should be used to see your sequencing runs distribution.
- also takes the bam file convert them to the fastq and estimates and plots the visualization.
- BAM support added, it converts the BAM file, reads the fastq and also puts the fasta for sending to the user.
- Make sure that the BAM and the .pbi files are in the same folder and it will read all the BAM files for the run as a regular expression. 
- it uses [Pterm](https://github.com/pterm/pterm) to directly give you the visualization plots.

```
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/golang-pacbiohifi-visualization ±main⚡ » \
go run main.go -h
This pacbiohifi application provides the binning intervals of the sequence classification from the sequencing reads

Usage:
  flag [command]

Available Commands:
  bam
  completion  Generate the autocompletion script for the specified shell
  fastq
  help        Help about any command

Flags:
  -h, --help   help for flag

Use "flag [command] --help" for more information about a command.
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/golang-pacbiohifi-visualization ±main⚡ » \
go run main.go fastq -h
Analyzes the fastq files

Usage:
  flag fastq [flags]

Flags:
  -h, --help               help for fastq
  -i, --inputfile string   pacbiohifi reads file (default "inputfile for the fastq reads")
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/golang-pacbiohifi-visualization ±main⚡ » \
go run main.go bam -h
Analyzes the bam files

Usage:
  flag bam [flags]

Flags:
  -b, --bamfile string   pacbiohifi bam file or the folder containing the bam files along with the pbi files (default "inputfile should be bam")
  -h, --help             help for bam
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/golang-pacbiohifi-visualization ±main⚡ » \
go run main.go fastq -i ./sample-files/samplepacbiohifi.fastq

```

![](https://github.com/codecreatede/go-pacbiohifi-visualization/blob/main/go-pacbiohifivisualization.png)

Gaurav Sablok
