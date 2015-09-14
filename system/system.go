package system

import (
	"github.com/aymerick/charette/core"
	"github.com/aymerick/charette/rom"
)

// System represents a gaming system found in no-intro archives
type System struct {
	// system informations
	Infos Infos

	// options
	Options *core.Options

	// found games
	Games map[string]*rom.Game

	// skipped files
	Skipped []string
}

// New instanciates a new System
func New(infos Infos, options *core.Options) *System {
	return &System{
		Infos:   infos,
		Options: options,
		Games:   map[string]*rom.Game{},
	}
}

//
// var (
// 	// allowed file extensions
// 	extensions []string
// )
//
// func init() {
// 	// no-intro files extensions
// 	extensions = []string{
// 		".zip", ".7z", // compressed
// 		".a26",        // atari2600
// 		".fds",        // famicomdisksystem
// 		".gb", ".gbc", // gameboy
// 		".gba",        // gameboy advance
// 		".gg",         // gamegear
// 		".ms",         // mastersystem
// 		".gen", ".md", // megadrive
// 		".rom", ".msx1", ".msx2", // msx
// 		".nes",                 // nes
// 		".n64",                 // nintendo64
// 		".pce",                 // pcengine
// 		".iso", ".img", ".bin", // playstation
// 		".32x",         // sega32x
// 		".cue",         // segacd
// 		".sg",          // sg1000
// 		".smc", ".sfc", // snes
// 		".vb", // virtual boy
// 	}
// }
//
// // Run detects roms in directory and filters them
// func (s *System) Run() error {
// 	if s.Options.Debug {
// 		log.Printf("Scanning files: %s", s.Dir)
// 	}
//
// 	// process files
// 	infos, err := ioutil.ReadDir(s.Dir)
// 	if err != nil {
// 		return err
// 	}
//
// 	nb := 0
//
// 	for _, info := range infos {
// 		if info.IsDir() {
// 			// Skip directories
// 			continue
// 		}
//
// 		if !expectedExt(info) {
// 			continue
// 		}
//
// 		// process file
// 		if err := s.processFile(info); err != nil {
// 			fmt.Printf("ERR: %v\n", err)
// 		}
//
// 		nb++
// 	}
//
// 	fmt.Printf("Processed %v files\n", nb)
// 	fmt.Printf("Skipped %v files\n", len(s.Skipped))
// 	fmt.Printf("Found %v games\n", len(s.Games))
//
// 	if s.Options.Unzip {
// 		// unzip roms
// 		s.unzip()
// 	}
//
// 	return nil
// }
//
// // unzip decompress all selected roms
// func (s *System) unzip() {
// 	nb := 0
//
// 	fmt.Printf("Unzipping roms, please wait...\n")
//
// 	for _, g := range s.Games {
// 		r := g.BestRom()
// 		if r.IsZipped() {
// 			if err := s.unzipFile(r.Filename); err != nil {
// 				fmt.Printf("ERR: %v\n", err)
// 			} else {
// 				nb++
//
// 				if err := s.deleteFile(r.Filename); err != nil {
// 					fmt.Printf("ERR: %v\n", err)
// 				}
// 			}
// 		}
// 	}
//
// 	if nb > 0 {
// 		fmt.Printf("Unzipped %v files\n", nb)
// 	}
// }
//
// // unzipFile decompress given file
// func (s *System) unzipFile(fileName string) error {
// 	filePath := path.Join(s.Dir, fileName)
//
// 	zipReader, err := zip.OpenReader(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer zipReader.Close()
//
// 	for _, f := range zipReader.Reader.File {
// 		zipFile, err := f.Open()
// 		if err != nil {
// 			return err
// 		}
// 		defer zipFile.Close()
//
// 		path := filepath.Join(s.Dir, f.Name)
//
// 		if f.FileInfo().IsDir() {
// 			if !s.Options.Noop {
// 				os.MkdirAll(path, f.Mode())
// 			}
// 		} else {
// 			writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, f.Mode())
// 			if err != nil {
// 				return err
// 			}
// 			defer writer.Close()
//
// 			if !s.Options.Noop {
// 				if _, err = io.Copy(writer, zipFile); err != nil {
// 					return err
// 				}
// 			}
// 		}
// 	}
//
// 	return nil
// }
//
// // deleteFile deletes given file
// func (s *System) deleteFile(fileName string) error {
// 	filePath := path.Join(s.Dir, fileName)
//
// 	if !s.Options.Noop {
// 		if err := os.Remove(filePath); err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }
//
// // moveFile moves given file to garbage
// func (s *System) moveFile(fileName string) error {
// 	oldPath := path.Join(s.Dir, fileName)
// 	newPath := path.Join(s.Garbage, fileName)
//
// 	if s.Options.Debug {
// 		log.Printf("MOVING %s => %s", oldPath, newPath)
// 	}
//
// 	if s.Options.Noop {
// 		// NOOP
// 		return nil
// 	}
//
// 	return os.Rename(oldPath, newPath)
// }
//
// // processFile handles a new file
// func (s *System) processFile(info os.FileInfo) error {
// 	r := rom.New(info.Name())
// 	if err := r.Fill(); err != nil {
// 		return err
// 	}
//
// 	if skip, msg := s.skip(r); skip {
// 		if s.Options.Verbose {
// 			fmt.Printf("Skipped '%s': %s\n", r.Filename, msg)
// 		}
//
// 		s.Skipped = append(s.Skipped, r.Filename)
//
// 		return nil
// 	}
//
// 	if s.Games[r.Name] == nil {
// 		// it's a new game
// 		s.Games[r.Name] = rom.NewGame()
// 	}
//
// 	s.Games[r.Name].AddRom(r)
//
// 	return nil
// }
//
// // skip returns true if given rom must be skiped, with an explanation message
// func (s *System) skip(r *rom.Rom) (bool, string) {
// 	if s.Options.LeaveMeAlone && !r.HaveRegion(s.Options.Regions) {
// 		return true, fmt.Sprintf("Leave me alone: %v", r.Regions)
// 	}
//
// 	if r.Proto && !s.Options.KeepProto {
// 		return true, "Ignore proto"
// 	}
//
// 	if r.Beta && !s.Options.KeepBeta {
// 		return true, "Ignore beta"
// 	}
//
// 	if r.Bios {
// 		return true, "Ignore bios"
// 	}
//
// 	if r.Sample && !s.Options.KeepSample {
// 		return true, "Ignore sample"
// 	}
//
// 	if r.Demo && !s.Options.KeepDemo {
// 		return true, "Ignore demo"
// 	}
//
// 	if r.Pirate && !s.Options.KeepPirate {
// 		return true, "Ignore pirate"
// 	}
//
// 	if r.Promo && !s.Options.KeepPromo {
// 		return true, "Ignore promo"
// 	}
//
// 	return false, ""
// }
//
// // expectedExt returns true if this is an expected no-intro rom file extension
// func expectedExt(info os.FileInfo) bool {
// 	fileExt := filepath.Ext(info.Name())
//
// 	for _, ext := range extensions {
// 		if fileExt == ext {
// 			return true
// 		}
// 	}
//
// 	return false
// }
