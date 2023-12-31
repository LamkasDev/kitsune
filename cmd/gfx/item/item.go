package item

import (
	"github.com/LamkasDev/kurin/cmd/gameplay"
	"github.com/LamkasDev/kurin/cmd/gfx"
	"github.com/LamkasDev/kurin/cmd/gfx/render"
	"github.com/veandco/go-sdl2/sdl"
)

func GetKurinItemRect(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame, item *gameplay.KurinItem) sdl.Rect {
	graphic := layer.Data.(KurinRendererLayerItemData).Items[item.Type]
	return render.WorldToScreenRect(renderer, sdl.FRect{
		X: float32(item.Position.Base.X) - 0.5, Y: float32(item.Position.Base.Y) - 0.5,
		W: float32(graphic.Texture.Base.Size.W), H: float32(graphic.Texture.Base.Size.H),
	})
}

func RenderKurinItem(renderer *gfx.KurinRenderer, layer *gfx.KurinRendererLayer, game *gameplay.KurinGame, item *gameplay.KurinItem) *error {
	graphic := layer.Data.(KurinRendererLayerItemData).Items[item.Type]
	rect := GetKurinItemRect(renderer, layer, game, item)

	if game.HoveredItem == item && graphic.Outline != nil {
		if err := renderer.Renderer.Copy(graphic.Outline.Texture, nil, &rect); err != nil {
			return &err
		}
	}
	if err := renderer.Renderer.Copy(graphic.Texture.Base.Texture, nil, &rect); err != nil {
		return &err
	}

	return nil
}
