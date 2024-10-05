package main

/*
Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-25

A golang streaming application for visualization of the pacbiohifi reads coming from the pacbiohifi sequencing.
It takes both the fastq as well as the bam file and converts the reads to fastq and gives you the analysis.
You should have pbtk installed in your path and it will convert the fastq to fasta and then will plot the reads.

*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	inputfile string
	bamfile   string
)

var rootCmd = &cobra.Command{
	Use:  "flag",
	Long: "This pacbiohifi application provides the binning intervals of the sequence classification from the sequencing reads",
}

var fastqCmd = &cobra.Command{
	Use:  "fastq",
	Long: "Analyzes the fastq files",
	Run:  fastqFunc,
}

var bamCmd = &cobra.Command{
	Use:  "bam",
	Long: "Analyzes the bam files",
	Run:  bamFunc,
}

func init() {
	fastqCmd.Flags().
		StringVarP(&inputfile, "inputfile", "i", "inputfile for the fastq reads", "pacbiohifi reads file")

	bamCmd.Flags().
		StringVarP(&bamfile, "bamfile", "b", "inputfile should be bam", "pacbiohifi bam file or the folder containing the bam files along with the pbi files")

	rootCmd.AddCommand(fastqCmd)
	rootCmd.AddCommand(bamCmd)
}

func fastqFunc(cmd *cobra.Command, args []string) {
	type headerID struct {
		id string
	}
	type sequenceID struct {
		seq string
	}

	header := []headerID{}
	sequence := []sequenceID{}
	sequence5K := []string{}
	sequence10K := []string{}
	sequence20K := []string{}
	sequence30K := []string{}
	sequence40K := []string{}
	sequence50K := []string{}
	sequenceL50K := []string{}
	sequence100K := []string{}
	sequence110K := []string{}
	sequence120K := []string{}
	sequence130K := []string{}
	sequence140K := []string{}
	sequence150K := []string{}
	sequence160K := []string{}
	sequence170K := []string{}
	sequence180K := []string{}
	sequence190K := []string{}
	sequence200K := []string{}
	sequenceabove200K := []string{}

	fOpen, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), "@") {
			header = append(header, headerID{
				id: string(line),
			})
		}
		if strings.HasPrefix(string(line), "A") || strings.HasPrefix(string(line), "T") ||
			strings.HasPrefix(string(line), "G") ||
			strings.HasPrefix(string(line), "C") && !strings.Contains(string(line), ";") {
			sequence = append(sequence, sequenceID{
				seq: string(line),
			})
		}
	}

	for i := range sequence {
		if len(string(sequence[i].seq)) >= 1000 && len(string(sequence[i].seq)) <= 5000 {
			sequence5K = append(sequence5K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 5000 && len(string(sequence[i].seq)) <= 10000 {
			sequence10K = append(sequence10K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 10000 && len(string(sequence[i].seq)) <= 20000 {
			sequence20K = append(sequence20K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 20000 && len(string(sequence[i].seq)) <= 30000 {
			sequence30K = append(sequence30K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 30000 && len(string(sequence[i].seq)) <= 40000 {
			sequence40K = append(sequence40K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 40000 && len(string(sequence[i].seq)) <= 50000 {
			sequenceL50K = append(sequenceL50K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) >= 50000 && len(string(sequence[i].seq)) <= 100000 {
			sequence50K = append(sequence50K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) >= 100000 {
			sequence100K = append(sequence100K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 100000 && len(string(sequence[i].seq)) <= 110000 {
			sequence110K = append(sequence110K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 110000 && len(string(sequence[i].seq)) <= 120000 {
			sequence120K = append(sequence120K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 120000 && len(string(sequence[i].seq)) <= 130000 {
			sequence130K = append(sequence130K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 130000 && len(string(sequence[i].seq)) <= 140000 {
			sequence140K = append(sequence140K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 140000 && len(string(sequence[i].seq)) <= 150000 {
			sequence150K = append(sequence150K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 150000 && len(string(sequence[i].seq)) <= 160000 {
			sequence160K = append(sequence160K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 160000 && len(string(sequence[i].seq)) <= 170000 {
			sequence170K = append(sequence170K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 170000 && len(string(sequence[i].seq)) <= 180000 {
			sequence180K = append(sequence180K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 180000 && len(string(sequence[i].seq)) <= 190000 {
			sequence190K = append(sequence190K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 190000 && len(string(sequence[i].seq)) <= 200000 {
			sequence200K = append(sequence200K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 200000 {
			sequenceabove200K = append(sequenceabove200K, sequence[i].seq)
		}
	}

	bars := []pterm.Bar{
		{Label: "sequence5K", Value: len(sequence5K)},
		{Label: "sequence10K", Value: len(sequence10K)},
		{Label: "sequence20K", Value: len(sequence20K)},
		{Label: "sequence30K", Value: len(sequence30K)},
		{Label: "sequence40K", Value: len(sequence40K)},
		{Label: "sequenceL50K", Value: len(sequenceL50K)},
		{Label: "sequence50K", Value: len(sequence50K)},
		{Label: "sequence100K", Value: len(sequence100K)},
		{Label: "sequence110K", Value: len(sequence110K)},
		{Label: "sequence120K", Value: len(sequence120K)},
		{Label: "sequence130K", Value: len(sequence130K)},
		{Label: "sequence140K", Value: len(sequence140K)},
		{Label: "sequence150K", Value: len(sequence150K)},
		{Label: "sequence160K", Value: len(sequence160K)},
		{Label: "sequence170K", Value: len(sequence170K)},
		{Label: "sequence180K", Value: len(sequence180K)},
		{Label: "sequence190K", Value: len(sequence190K)},
		{Label: "sequence200K", Value: len(sequence200K)},
		{Label: "sequenceabove200K", Value: len(sequenceabove200K)},
	}
	pterm.Info.Println("Profiling your Pacbiohifi sequencing reads")
	pterm.DefaultBarChart.WithBars(bars).Render()
}

func bamFunc(cmd *cobra.Command, args []string) {
	fastq, err := exec.Command("bam2fastq", "-o", "pacbio", "*.bam").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(fastq))
	fasta, err := exec.Command("bam2fasta", "-o", "pacbio", "*.bam").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fasta)
	out, err := exec.Command("gunzip", "pacbio.fastq.gz").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
	type headerID struct {
		id string
	}
	type sequenceID struct {
		seq string
	}

	header := []headerID{}
	sequence := []sequenceID{}
	sequence5K := []string{}
	sequence10K := []string{}
	sequence20K := []string{}
	sequence30K := []string{}
	sequence40K := []string{}
	sequence50K := []string{}
	sequenceL50K := []string{}
	sequence100K := []string{}
	sequence110K := []string{}
	sequence120K := []string{}
	sequence130K := []string{}
	sequence140K := []string{}
	sequence150K := []string{}
	sequence160K := []string{}
	sequence170K := []string{}
	sequence180K := []string{}
	sequence190K := []string{}
	sequence200K := []string{}
	sequenceabove200K := []string{}

	fOpen, err := os.Open("pacbio.fastq")
	if err != nil {
		log.Fatal(err)
	}
	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), "@") {
			header = append(header, headerID{
				id: string(line),
			})
		}
		if strings.HasPrefix(string(line), "A") || strings.HasPrefix(string(line), "T") ||
			strings.HasPrefix(string(line), "G") ||
			strings.HasPrefix(string(line), "C") && !strings.Contains(string(line), ";") {
			sequence = append(sequence, sequenceID{
				seq: string(line),
			})
		}
	}

	for i := range sequence {
		if len(string(sequence[i].seq)) >= 1000 && len(string(sequence[i].seq)) <= 5000 {
			sequence5K = append(sequence5K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 5000 && len(string(sequence[i].seq)) <= 10000 {
			sequence10K = append(sequence10K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 10000 && len(string(sequence[i].seq)) <= 20000 {
			sequence20K = append(sequence20K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 20000 && len(string(sequence[i].seq)) <= 30000 {
			sequence30K = append(sequence30K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 30000 && len(string(sequence[i].seq)) <= 40000 {
			sequence40K = append(sequence40K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 40000 && len(string(sequence[i].seq)) <= 50000 {
			sequenceL50K = append(sequenceL50K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) >= 50000 && len(string(sequence[i].seq)) <= 100000 {
			sequence50K = append(sequence50K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) >= 100000 {
			sequence100K = append(sequence100K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 100000 && len(string(sequence[i].seq)) <= 110000 {
			sequence110K = append(sequence110K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 110000 && len(string(sequence[i].seq)) <= 120000 {
			sequence120K = append(sequence120K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 120000 && len(string(sequence[i].seq)) <= 130000 {
			sequence130K = append(sequence130K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 130000 && len(string(sequence[i].seq)) <= 140000 {
			sequence140K = append(sequence140K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 140000 && len(string(sequence[i].seq)) <= 150000 {
			sequence150K = append(sequence150K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 150000 && len(string(sequence[i].seq)) <= 160000 {
			sequence160K = append(sequence160K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 160000 && len(string(sequence[i].seq)) <= 170000 {
			sequence170K = append(sequence170K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 170000 && len(string(sequence[i].seq)) <= 180000 {
			sequence180K = append(sequence180K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 180000 && len(string(sequence[i].seq)) <= 190000 {
			sequence190K = append(sequence190K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 190000 && len(string(sequence[i].seq)) <= 200000 {
			sequence200K = append(sequence200K, sequence[i].seq)
		}
		if len(string(sequence[i].seq)) > 200000 {
			sequenceabove200K = append(sequenceabove200K, sequence[i].seq)
		}

		bars := []pterm.Bar{
			{Label: "sequence5K", Value: len(sequence5K)},
			{Label: "sequence10K", Value: len(sequence10K)},
			{Label: "sequence20K", Value: len(sequence20K)},
			{Label: "sequence30K", Value: len(sequence30K)},
			{Label: "sequence40K", Value: len(sequence40K)},
			{Label: "sequenceL50K", Value: len(sequenceL50K)},
			{Label: "sequence50K", Value: len(sequence50K)},
			{Label: "sequence100K", Value: len(sequence100K)},
			{Label: "sequence110K", Value: len(sequence110K)},
			{Label: "sequence120K", Value: len(sequence120K)},
			{Label: "sequence130K", Value: len(sequence130K)},
			{Label: "sequence140K", Value: len(sequence140K)},
			{Label: "sequence150K", Value: len(sequence150K)},
			{Label: "sequence160K", Value: len(sequence160K)},
			{Label: "sequence170K", Value: len(sequence170K)},
			{Label: "sequence180K", Value: len(sequence180K)},
			{Label: "sequence190K", Value: len(sequence190K)},
			{Label: "sequence200K", Value: len(sequence200K)},
			{Label: "sequenceabove200K", Value: len(sequenceabove200K)},
		}
		pterm.Info.Println("Profiling your Pacbiohifi sequencing reads")
		pterm.DefaultBarChart.WithBars(bars).Render()

	}
}
