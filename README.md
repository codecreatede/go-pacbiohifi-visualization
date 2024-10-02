go-pacbiohifi-visualization

- bins the sequenced pacbiohifi reads for visualization
- this should be used to see your sequencing runs distribution.
- also takes the bam file convert them to the fastq and estimates and plots the visualization.
- BAM support added, it converts the BAM file, reads the fastq and also puts the fasta for sending to the user.
- Make sure that the BAM and the .pbi files are in the same folder and it will read all the BAM files for the run as a regular expression. 
- it uses [Pterm](https://github.com/pterm/pterm) to directly give you the visualization plots.

```
gauravsablok@gaurav-sablok ~/Desktop/codecreatede/golang/golang-pacbiohifi-visualization ±main⚡ » \
go run main.go fastq -h
Analyzes the fastq files

Usage:
  flag fastq [flags]

Flags:
  -h, --help               help for fastq
  -i, --inputfile string   pacbiohifi reads file (default "inputfile for the fastq reads")
gauravsablok@gaurav-sablok ~/Desktop/codecreatede/golang/golang-pacbiohifi-visualization ±main⚡ » \
go run main.go bam -h
Analyzes the bam files

Usage:
  flag bam [flags]

Flags:
  -b, --bamfile string   pacbiohifi bam file or the folder containing the bam files along with the pbi files (default "inputfile should be bam")
  -h, --help             help for bam


```

Gaurav Sablok
