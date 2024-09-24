go-pacbiohifi-visualization

- bins the sequenced pacbiohifi reads for visualization
- this should be used to see your sequencing runs distribution.
- also takes the bam file convert them to the fastq and estimates and plots the visualization.
- it uses [Pterm](https://github.com/pterm/pterm) to directly give you the visualization plots.
```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/golang-pacbiohifi-visualization% \
go run main.go -h
This pacbiohifi application provides the binning intervals of the sequence classification from the sequencing reads

Usage:
  flag [flags]

Flags:
  -b, --bamfile string     pacbiohifi bam file (default "inputfile should be bam")
  -h, --help               help for flag
  -i, --inputfile string   pacbiohifi reads file (default "inputfile for the fastq reads")
```

Gaurav Sablok
