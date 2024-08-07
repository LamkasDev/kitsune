package keybinds

import (
	"encoding/json"
	"os"
	"path"

	"github.com/LamkasDev/kurin/cmd/common/constants"
	"github.com/LamkasDev/kurin/cmd/common/sdlutils"
	"github.com/LamkasDev/kurin/cmd/event"
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

type KurinEventLayerKeybindsData struct {
}

func NewKurinEventLayerKeybinds() *event.KurinEventLayer {
	return &event.KurinEventLayer{
		Load:    LoadKurinEventLayerKeybinds,
		Process: ProcessKurinEventLayerKeybinds,
		Data:    KurinEventLayerKeybindsData{},
	}
}

func LoadKurinEventLayerKeybinds(manager *event.KurinEventManager, layer *event.KurinEventLayer) *error {
	return nil
}

func ProcessKurinEventLayerKeybinds(manager *event.KurinEventManager, layer *event.KurinEventLayer, game *gameplay.KurinGame) *error {
	if manager.Keyboard.Pending != nil {
		switch *manager.Keyboard.Pending {
		case sdl.K_x:
			if game.SelectedCharacter == nil {
				return nil
			}
			switch game.SelectedCharacter.ActiveHand {
			case gameplay.KurinHandLeft:
				game.SelectedCharacter.ActiveHand = gameplay.KurinHandRight
			case gameplay.KurinHandRight:
				game.SelectedCharacter.ActiveHand = gameplay.KurinHandLeft
			}
		case sdl.K_q:
			if game.SelectedCharacter == nil {
				return nil
			}
			gameplay.DropKurinItemFromCharacter(game, game.SelectedCharacter)
		case sdl.K_r:
			if game.SelectedCharacter == nil {
				return nil
			}
			item := game.SelectedCharacter.Inventory.Hands[game.SelectedCharacter.ActiveHand]
			if !gameplay.DropKurinItemFromCharacter(game, game.SelectedCharacter) {
				return nil
			}
			wpos := render.ScreenToWorldPosition(manager.Renderer, manager.Renderer.Context.MousePosition)
			force := gameplay.KurinForce{
				Item: item,
				Target: sdlutils.PointToFPointCenter(wpos),
			}
			game.ForceController.Forces[item] = &force
		case sdl.K_f:
			switch manager.Renderer.Context.CameraMode {
			case gfx.KurinRendererCameraModeCharacter:
				manager.Renderer.Context.CameraMode = gfx.KurinRendererCameraModeFree
				game.SelectedCharacter = nil
			case gfx.KurinRendererCameraModeFree:
				manager.Renderer.Context.CameraMode = gfx.KurinRendererCameraModeCharacter
				game.SelectedCharacter = game.Characters[0]
			}
		case sdl.K_s:
			if manager.Keyboard.Pressed[sdl.K_LCTRL] {
				gameJson, err := json.Marshal(game)
				if err != nil {
					return &err
				}
				os.WriteFile(path.Join(constants.TempSavesPath, "save.json"), gameJson, 777)
			}
		default:
			return nil
		}
		manager.Keyboard.Pending = nil
	}

	return nil
}
