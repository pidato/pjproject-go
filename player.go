package pj

import (
	"github.com/pidato/pjproject-go/pjsua2"
	"sync"
)

type Player struct {
	*Audio
	player pjsua2.AudioMediaPlayer

	mu sync.Mutex
}

func NewPlayer(file string) (player *Player, err error) {
	player = &Player{}
	(func() {
		defer func() {
			if e := recover(); e != nil {
				er, ok := e.(error)
				if ok {
					err = er
				}
			}
		}()
		player.player = pjsua2.NewDirectorAudioMediaPlayer(player)
		player.player.CreatePlayer(file, pjsua2.PJMEDIA_FILE_NO_LOOP)
	})()

	if err != nil {
		player.Audio = NewAudio(player.player)
	}

	player.player.SetPos(0)

	return player, err
}

func (p *Player) OnEof() (_swig_ret bool) {
	_swig_ret = true
	return _swig_ret
}
