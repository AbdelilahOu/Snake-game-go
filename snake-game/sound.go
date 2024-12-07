package SnakeGame

import (
	"bytes"
	_ "embed"
	"log"

	Resources "github.com/AbdelilahOu/Snake-game-go/resources"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

// AudioManager handles loading, playing, and managing game audio
type AudioManager struct {
	audioContext *audio.Context
	sounds       map[string]*audio.Player
	music        map[string]*audio.Player
	currentMusic *audio.Player
}

// NewAudioManager creates a new AudioManager
func NewAudioManager() *AudioManager {
	return &AudioManager{
		audioContext: audio.NewContext(44100),
		sounds:       make(map[string]*audio.Player),
		music:        make(map[string]*audio.Player),
	}
}

var GlobalAudioManager *AudioManager

func init() {
	GlobalAudioManager = NewAudioManager()
	GlobalAudioManager.LoadSoundFile("game-over", Resources.GameOverSound)
	GlobalAudioManager.LoadSoundFile("food", Resources.FoodSound)
	GlobalAudioManager.LoadSoundFile("move", Resources.MoveSound)
	GlobalAudioManager.LoadMusicFile("music", Resources.MusicSound)
}

// LoadSoundFile loads an embedded MP3 sound file
func (am *AudioManager) LoadSoundFile(name string, soundFile []byte) error {
	// Decode the MP3 file
	decoded, err := mp3.Decode(am.audioContext, bytes.NewReader(soundFile))
	if err != nil {
		return err
	}

	// Create an audio player
	player, err := am.audioContext.NewPlayer(decoded)
	if err != nil {
		return err
	}

	// Store the player in the sounds map
	am.sounds[name] = player
	return nil
}

// LoadMusicFile loads an embedded MP3 music file
func (am *AudioManager) LoadMusicFile(name string, musicFile []byte) error {
	// Decode the MP3 file
	decoded, err := mp3.Decode(am.audioContext, bytes.NewReader(musicFile))
	if err != nil {
		return err
	}

	// Create an audio player
	player, err := am.audioContext.NewPlayer(decoded)
	if err != nil {
		return err
	}

	// Store the player in the music map
	am.music[name] = player
	return nil
}

// PlaySound plays a sound effect
func (am *AudioManager) PlaySound(name string) {
	player, exists := am.sounds[name]
	if !exists {
		log.Printf("Sound %s not found", name)
		return
	}

	// Rewind to start and play
	player.Rewind()
	player.Play()
}

// PlayMusic starts playing background music
func (am *AudioManager) PlayMusic(name string) {
	// Stop current music if playing
	if am.currentMusic != nil {
		am.currentMusic.Pause()
	}

	player, exists := am.music[name]
	if !exists {
		log.Printf("Music %s not found", name)
		return
	}

	// Set as current music, rewind, and play
	am.currentMusic = player
	player.Rewind()
	player.SetVolume(0.2) // Default volume
	player.Play()
}

// StopMusic stops the current background music
func (am *AudioManager) StopMusic() {
	if am.currentMusic != nil {
		am.currentMusic.Pause()
		am.currentMusic.Rewind()
	}
}

// SetMusicVolume adjusts the volume of the current music (0.0 to 1.0)
func (am *AudioManager) SetMusicVolume(volume float64) {
	if am.currentMusic != nil {
		am.currentMusic.SetVolume(volume)
	}
}

// Close releases all audio resources
func (am *AudioManager) Close() {
	// Stop and close all sound players
	for _, player := range am.sounds {
		player.Close()
	}

	// Stop and close all music players
	for _, player := range am.music {
		player.Close()
	}
}
