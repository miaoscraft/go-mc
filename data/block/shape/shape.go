// Package shape stores information about the shapes of blocks in Minecraft.
package shape

import (
	"github.com/Tnze/go-mc/data/block"
)

// ID describes a numeric shape ID.
type ID uint32

// Shape describes how collisions should be tested for an object.
type Shape struct {
	ID    ID
	Boxes []BoundingBox
}

type BoundingTriplet struct {
	X, Y, Z float64
}

type BoundingBox struct {
	Min, Max BoundingTriplet
}

// ByBlockID is an index of shapes for each minecraft block variant.
var ByBlockID = map[block.ID][]ID{
	block.CyanWallBanner.ID:                  []ID{0},
	block.Diorite.ID:                         []ID{1},
	block.LightBlueBed.ID:                    []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.BrownBed.ID:                        []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.DarkOakWallSign.ID:                 []ID{0},
	block.JunglePressurePlate.ID:             []ID{0},
	block.SandstoneStairs.ID:                 []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.PinkTerracotta.ID:                  []ID{1},
	block.Air.ID:                             []ID{0},
	block.DiamondBlock.ID:                    []ID{1},
	block.SpruceTrapdoor.ID:                  []ID{54, 54, 54, 54, 89, 89, 89, 89, 54, 54, 54, 54, 88, 88, 88, 88, 57, 57, 57, 57, 89, 89, 89, 89, 57, 57, 57, 57, 88, 88, 88, 88, 56, 56, 56, 56, 89, 89, 89, 89, 56, 56, 56, 56, 88, 88, 88, 88, 55, 55, 55, 55, 89, 89, 89, 89, 55, 55, 55, 55, 88, 88, 88, 88},
	block.RedstoneBlock.ID:                   []ID{1},
	block.ChiseledRedSandstone.ID:            []ID{1},
	block.WhiteShulkerBox.ID:                 []ID{1},
	block.TwistingVinesPlant.ID:              []ID{0},
	block.StrippedSpruceWood.ID:              []ID{1},
	block.BrownMushroomBlock.ID:              []ID{1},
	block.GreenWallBanner.ID:                 []ID{0},
	block.QuartzSlab.ID:                      []ID{169, 169, 61, 61, 1, 1},
	block.SpruceSapling.ID:                   []ID{0},
	block.Furnace.ID:                         []ID{1},
	block.HeavyWeightedPressurePlate.ID:      []ID{0},
	block.CyanStainedGlassPane.ID:            []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.PrismarineStairs.ID:                []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.SeaLantern.ID:                      []ID{1},
	block.DeadBrainCoral.ID:                  []ID{0},
	block.JungleSapling.ID:                   []ID{0},
	block.StrippedOakLog.ID:                  []ID{1},
	block.AcaciaLeaves.ID:                    []ID{1},
	block.OrangeWool.ID:                      []ID{1},
	block.Bookshelf.ID:                       []ID{1},
	block.WhiteStainedGlassPane.ID:           []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.LightBlueConcrete.ID:               []ID{1},
	block.SmoothSandstoneSlab.ID:             []ID{169, 169, 61, 61, 1, 1},
	block.DioriteSlab.ID:                     []ID{169, 169, 61, 61, 1, 1},
	block.Bedrock.ID:                         []ID{1},
	block.OakLeaves.ID:                       []ID{1},
	block.WitherRose.ID:                      []ID{0},
	block.EmeraldOre.ID:                      []ID{1},
	block.LightBlueWallBanner.ID:             []ID{0},
	block.LightGrayConcretePowder.ID:         []ID{1},
	block.HornCoralWallFan.ID:                []ID{0},
	block.NetherWartBlock.ID:                 []ID{1},
	block.StrippedSpruceLog.ID:               []ID{1},
	block.DeadBush.ID:                        []ID{0},
	block.Ice.ID:                             []ID{1},
	block.LightGrayStainedGlass.ID:           []ID{1},
	block.InfestedStoneBricks.ID:             []ID{1},
	block.BrickStairs.ID:                     []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.PurpurPillar.ID:                    []ID{1},
	block.LimeConcretePowder.ID:              []ID{1},
	block.TubeCoralWallFan.ID:                []ID{0},
	block.BlastFurnace.ID:                    []ID{1},
	block.CrimsonFungus.ID:                   []ID{0},
	block.Cobblestone.ID:                     []ID{1},
	block.JackOLantern.ID:                    []ID{1},
	block.BlackTerracotta.ID:                 []ID{1},
	block.TwistingVines.ID:                   []ID{0},
	block.Netherrack.ID:                      []ID{1},
	block.InfestedChiseledStoneBricks.ID:     []ID{1},
	block.NetherBrickStairs.ID:               []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.CrimsonSign.ID:                     []ID{0},
	block.OakSapling.ID:                      []ID{0},
	block.AcaciaWallSign.ID:                  []ID{0},
	block.LightBlueTerracotta.ID:             []ID{1},
	block.GrayTerracotta.ID:                  []ID{1},
	block.DarkOakStairs.ID:                   []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.PrismarineBrickStairs.ID:           []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.RedBanner.ID:                       []ID{0},
	block.PolishedDiorite.ID:                 []ID{1},
	block.PumpkinStem.ID:                     []ID{0},
	block.PurpleTerracotta.ID:                []ID{1},
	block.GrayWallBanner.ID:                  []ID{0},
	block.YellowShulkerBox.ID:                []ID{1},
	block.CartographyTable.ID:                []ID{1},
	block.YellowCarpet.ID:                    []ID{170},
	block.DeadBubbleCoralWallFan.ID:          []ID{0},
	block.NetherBrickWall.ID:                 []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.WarpedWallSign.ID:                  []ID{0},
	block.PottedRedTulip.ID:                  []ID{156},
	block.JungleFenceGate.ID:                 []ID{0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75},
	block.GreenConcretePowder.ID:             []ID{1},
	block.DeadTubeCoralBlock.ID:              []ID{1},
	block.WarpedSign.ID:                      []ID{0},
	block.PolishedBlackstoneBrickWall.ID:     []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.MagentaBanner.ID:                   []ID{0},
	block.SmoothStone.ID:                     []ID{1},
	block.EndGateway.ID:                      []ID{0},
	block.OrangeShulkerBox.ID:                []ID{1},
	block.GrayShulkerBox.ID:                  []ID{1},
	block.BubbleCoralBlock.ID:                []ID{1},
	block.Stone.ID:                           []ID{1},
	block.CyanBed.ID:                         []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.Tripwire.ID:                        []ID{0},
	block.DioriteStairs.ID:                   []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.GraniteWall.ID:                     []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.BlackstoneStairs.ID:                []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.KelpPlant.ID:                       []ID{0},
	block.IronBlock.ID:                       []ID{1},
	block.Tnt.ID:                             []ID{1},
	block.BirchSign.ID:                       []ID{0},
	block.SpruceStairs.ID:                    []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.BrownStainedGlassPane.ID:           []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.Barrier.ID:                         []ID{1},
	block.CutSandstoneSlab.ID:                []ID{169, 169, 61, 61, 1, 1},
	block.CaveAir.ID:                         []ID{0},
	block.CrimsonTrapdoor.ID:                 []ID{54, 54, 54, 54, 89, 89, 89, 89, 54, 54, 54, 54, 88, 88, 88, 88, 57, 57, 57, 57, 89, 89, 89, 89, 57, 57, 57, 57, 88, 88, 88, 88, 56, 56, 56, 56, 89, 89, 89, 89, 56, 56, 56, 56, 88, 88, 88, 88, 55, 55, 55, 55, 89, 89, 89, 89, 55, 55, 55, 55, 88, 88, 88, 88},
	block.SpruceLog.ID:                       []ID{1},
	block.SpruceSign.ID:                      []ID{0},
	block.StoneButton.ID:                     []ID{0},
	block.CarvedPumpkin.ID:                   []ID{1},
	block.OakButton.ID:                       []ID{0},
	block.IronTrapdoor.ID:                    []ID{54, 54, 54, 54, 89, 89, 89, 89, 54, 54, 54, 54, 88, 88, 88, 88, 57, 57, 57, 57, 89, 89, 89, 89, 57, 57, 57, 57, 88, 88, 88, 88, 56, 56, 56, 56, 89, 89, 89, 89, 56, 56, 56, 56, 88, 88, 88, 88, 55, 55, 55, 55, 89, 89, 89, 89, 55, 55, 55, 55, 88, 88, 88, 88},
	block.BlackConcretePowder.ID:             []ID{1},
	block.LapisOre.ID:                        []ID{1},
	block.StoneBrickWall.ID:                  []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.RedWool.ID:                         []ID{1},
	block.Torch.ID:                           []ID{0},
	block.EnchantingTable.ID:                 []ID{10},
	block.DarkOakButton.ID:                   []ID{0},
	block.StoneSlab.ID:                       []ID{169, 169, 61, 61, 1, 1},
	block.CyanConcrete.ID:                    []ID{1},
	block.PrismarineWall.ID:                  []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.WarpedNylium.ID:                    []ID{1},
	block.BirchPlanks.ID:                     []ID{1},
	block.PurpleWool.ID:                      []ID{1},
	block.Loom.ID:                            []ID{1},
	block.CrimsonSlab.ID:                     []ID{169, 169, 61, 61, 1, 1},
	block.PolishedBlackstoneButton.ID:        []ID{0},
	block.OrangeBed.ID:                       []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.Ladder.ID:                          []ID{54, 54, 57, 57, 56, 56, 55, 55},
	block.PinkStainedGlass.ID:                []ID{1},
	block.PetrifiedOakSlab.ID:                []ID{169, 169, 61, 61, 1, 1},
	block.BlackConcrete.ID:                   []ID{1},
	block.DeadBrainCoralFan.ID:               []ID{0},
	block.MossyStoneBrickWall.ID:             []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.TripwireHook.ID:                    []ID{0},
	block.JungleFence.ID:                     []ID{66, 67, 66, 67, 68, 69, 68, 69, 70, 71, 70, 71, 72, 73, 72, 73, 74, 75, 74, 75, 76, 77, 76, 77, 78, 79, 78, 79, 80, 65, 80, 65},
	block.DarkOakDoor.ID:                     []ID{55, 55, 54, 54, 56, 56, 54, 54, 55, 55, 54, 54, 56, 56, 54, 54, 56, 56, 57, 57, 55, 55, 57, 57, 56, 56, 57, 57, 55, 55, 57, 57, 54, 54, 56, 56, 57, 57, 56, 56, 54, 54, 56, 56, 57, 57, 56, 56, 57, 57, 55, 55, 54, 54, 55, 55, 57, 57, 55, 55, 54, 54, 55, 55},
	block.BrownShulkerBox.ID:                 []ID{1},
	block.OrangeConcretePowder.ID:            []ID{1},
	block.PurpleConcretePowder.ID:            []ID{1},
	block.Beacon.ID:                          []ID{1},
	block.LightGrayTerracotta.ID:             []ID{1},
	block.OrangeWallBanner.ID:                []ID{0},
	block.LightBlueShulkerBox.ID:             []ID{1},
	block.PinkWool.ID:                        []ID{1},
	block.Farmland.ID:                        []ID{53},
	block.JungleWallSign.ID:                  []ID{0},
	block.QuartzStairs.ID:                    []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.PackedIce.ID:                       []ID{1},
	block.LightBlueGlazedTerracotta.ID:       []ID{1},
	block.WarpedFence.ID:                     []ID{66, 67, 66, 67, 68, 69, 68, 69, 70, 71, 70, 71, 72, 73, 72, 73, 74, 75, 74, 75, 76, 77, 76, 77, 78, 79, 78, 79, 80, 65, 80, 65},
	block.YellowBed.ID:                       []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.DarkPrismarineStairs.ID:            []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.PinkWallBanner.ID:                  []ID{0},
	block.MagentaConcretePowder.ID:           []ID{1},
	block.WarpedStairs.ID:                    []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.PinkShulkerBox.ID:                  []ID{1},
	block.MagentaWool.ID:                     []ID{1},
	block.CraftingTable.ID:                   []ID{1},
	block.ZombieWallHead.ID:                  []ID{158, 159, 160, 161},
	block.Hopper.ID:                          []ID{164, 165, 166, 167, 168, 164, 165, 166, 167, 168},
	block.LightGrayBanner.ID:                 []ID{0},
	block.MagentaWallBanner.ID:               []ID{0},
	block.CobblestoneSlab.ID:                 []ID{169, 169, 61, 61, 1, 1},
	block.PolishedDioriteSlab.ID:             []ID{169, 169, 61, 61, 1, 1},
	block.GraniteSlab.ID:                     []ID{169, 169, 61, 61, 1, 1},
	block.StructureBlock.ID:                  []ID{1},
	block.SprucePlanks.ID:                    []ID{1},
	block.OakLog.ID:                          []ID{1},
	block.AcaciaWood.ID:                      []ID{1},
	block.Melon.ID:                           []ID{1},
	block.GreenGlazedTerracotta.ID:           []ID{1},
	block.DeadTubeCoral.ID:                   []ID{0},
	block.SandstoneWall.ID:                   []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.AzureBluet.ID:                      []ID{0},
	block.BirchWood.ID:                       []ID{1},
	block.LightGrayWool.ID:                   []ID{1},
	block.Cactus.ID:                          []ID{64},
	block.LimeTerracotta.ID:                  []ID{1},
	block.ChorusPlant.ID:                     []ID{175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 174},
	block.EndStoneBrickStairs.ID:             []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.PolishedBlackstoneBrickSlab.ID:     []ID{169, 169, 61, 61, 1, 1},
	block.PottedCactus.ID:                    []ID{156},
	block.Dropper.ID:                         []ID{1},
	block.FrostedIce.ID:                      []ID{1},
	block.RedGlazedTerracotta.ID:             []ID{1},
	block.BubbleCoralFan.ID:                  []ID{0},
	block.Scaffolding.ID:                     []ID{246},
	block.Jigsaw.ID:                          []ID{1},
	block.SnowBlock.ID:                       []ID{1},
	block.MagentaTerracotta.ID:               []ID{1},
	block.DeadHornCoralWallFan.ID:            []ID{0},
	block.PolishedBlackstoneSlab.ID:          []ID{169, 169, 61, 61, 1, 1},
	block.PottedOakSapling.ID:                []ID{156},
	block.PottedPinkTulip.ID:                 []ID{156},
	block.LimeStainedGlassPane.ID:            []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.SmoothQuartz.ID:                    []ID{1},
	block.PurpleGlazedTerracotta.ID:          []ID{1},
	block.SoulFire.ID:                        []ID{0},
	block.JungleLeaves.ID:                    []ID{1},
	block.QuartzBlock.ID:                     []ID{1},
	block.RedStainedGlassPane.ID:             []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.DarkPrismarineSlab.ID:              []ID{169, 169, 61, 61, 1, 1},
	block.OrangeGlazedTerracotta.ID:          []ID{1},
	block.StrippedBirchWood.ID:               []ID{1},
	block.Allium.ID:                          []ID{0},
	block.TrappedChest.ID:                    []ID{48, 48, 49, 49, 50, 50, 48, 48, 50, 50, 49, 49, 48, 48, 51, 51, 52, 52, 48, 48, 52, 52, 51, 51},
	block.OrangeConcrete.ID:                  []ID{1},
	block.ChiseledNetherBricks.ID:            []ID{1},
	block.AcaciaSapling.ID:                   []ID{0},
	block.DarkOakSapling.ID:                  []ID{0},
	block.IronDoor.ID:                        []ID{55, 55, 54, 54, 56, 56, 54, 54, 55, 55, 54, 54, 56, 56, 54, 54, 56, 56, 57, 57, 55, 55, 57, 57, 56, 56, 57, 57, 55, 55, 57, 57, 54, 54, 56, 56, 57, 57, 56, 56, 54, 54, 56, 56, 57, 57, 56, 56, 57, 57, 55, 55, 54, 54, 55, 55, 57, 57, 55, 55, 54, 54, 55, 55},
	block.AcaciaTrapdoor.ID:                  []ID{54, 54, 54, 54, 89, 89, 89, 89, 54, 54, 54, 54, 88, 88, 88, 88, 57, 57, 57, 57, 89, 89, 89, 89, 57, 57, 57, 57, 88, 88, 88, 88, 56, 56, 56, 56, 89, 89, 89, 89, 56, 56, 56, 56, 88, 88, 88, 88, 55, 55, 55, 55, 89, 89, 89, 89, 55, 55, 55, 55, 88, 88, 88, 88},
	block.InfestedStone.ID:                   []ID{1},
	block.FireCoral.ID:                       []ID{0},
	block.BrainCoralWallFan.ID:               []ID{0},
	block.BlueOrchid.ID:                      []ID{0},
	block.Pumpkin.ID:                         []ID{1},
	block.OakTrapdoor.ID:                     []ID{54, 54, 54, 54, 89, 89, 89, 89, 54, 54, 54, 54, 88, 88, 88, 88, 57, 57, 57, 57, 89, 89, 89, 89, 57, 57, 57, 57, 88, 88, 88, 88, 56, 56, 56, 56, 89, 89, 89, 89, 56, 56, 56, 56, 88, 88, 88, 88, 55, 55, 55, 55, 89, 89, 89, 89, 55, 55, 55, 55, 88, 88, 88, 88},
	block.OrangeCarpet.ID:                    []ID{170},
	block.BrownCarpet.ID:                     []ID{170},
	block.NetherBrickSlab.ID:                 []ID{169, 169, 61, 61, 1, 1},
	block.PottedAzureBluet.ID:                []ID{156},
	block.DriedKelpBlock.ID:                  []ID{1},
	block.CobblestoneStairs.ID:               []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.Dispenser.ID:                       []ID{1},
	block.Lever.ID:                           []ID{0},
	block.DeadHornCoralFan.ID:                []ID{0},
	block.PolishedDioriteStairs.ID:           []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.Grindstone.ID:                      []ID{248, 248, 249, 249, 247, 250, 251, 252, 253, 253, 254, 254},
	block.PolishedBasalt.ID:                  []ID{1},
	block.QuartzBricks.ID:                    []ID{1},
	block.Podzol.ID:                          []ID{1},
	block.StrippedAcaciaWood.ID:              []ID{1},
	block.MossyCobblestoneWall.ID:            []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.Target.ID:                          []ID{1},
	block.Fern.ID:                            []ID{0},
	block.TallSeagrass.ID:                    []ID{0},
	block.GreenCarpet.ID:                     []ID{170},
	block.SmoothSandstone.ID:                 []ID{1},
	block.BubbleCoralWallFan.ID:              []ID{0},
	block.Conduit.ID:                         []ID{244},
	block.CyanWool.ID:                        []ID{1},
	block.GreenTerracotta.ID:                 []ID{1},
	block.LargeFern.ID:                       []ID{0},
	block.RedNetherBricks.ID:                 []ID{1},
	block.WarpedSlab.ID:                      []ID{169, 169, 61, 61, 1, 1},
	block.CrimsonStairs.ID:                   []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.StrippedJungleWood.ID:              []ID{1},
	block.BirchStairs.ID:                     []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.CobblestoneWall.ID:                 []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.GrayConcrete.ID:                    []ID{1},
	block.YellowConcretePowder.ID:            []ID{1},
	block.TubeCoralBlock.ID:                  []ID{1},
	block.MossyCobblestoneSlab.ID:            []ID{169, 169, 61, 61, 1, 1},
	block.DarkOakPlanks.ID:                   []ID{1},
	block.Mycelium.ID:                        []ID{1},
	block.DeadBrainCoralWallFan.ID:           []ID{0},
	block.PinkBed.ID:                         []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.BrownConcrete.ID:                   []ID{1},
	block.CrimsonWallSign.ID:                 []ID{0},
	block.Gravel.ID:                          []ID{1},
	block.StrippedAcaciaLog.ID:               []ID{1},
	block.GrayBed.ID:                         []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.Beehive.ID:                         []ID{1},
	block.HoneycombBlock.ID:                  []ID{1},
	block.SmithingTable.ID:                   []ID{1},
	block.LimeBed.ID:                         []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.Cornflower.ID:                      []ID{0},
	block.DarkOakPressurePlate.ID:            []ID{0},
	block.BrewingStand.ID:                    []ID{107},
	block.EndStone.ID:                        []ID{1},
	block.DeadBubbleCoral.ID:                 []ID{0},
	block.DeadTubeCoralFan.ID:                []ID{0},
	block.BeeNest.ID:                         []ID{1},
	block.SoulWallTorch.ID:                   []ID{0},
	block.DragonHead.ID:                      []ID{157},
	block.Smoker.ID:                          []ID{1},
	block.PottedWarpedFungus.ID:              []ID{156},
	block.PottedWarpedRoots.ID:               []ID{156},
	block.CoalOre.ID:                         []ID{1},
	block.Grass.ID:                           []ID{0},
	block.SpruceWallSign.ID:                  []ID{0},
	block.PurpurStairs.ID:                    []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.MagentaGlazedTerracotta.ID:         []ID{1},
	block.BlueBed.ID:                         []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.OxeyeDaisy.ID:                      []ID{0},
	block.Beetroots.ID:                       []ID{0},
	block.Snow.ID:                            []ID{0, 58, 59, 60, 61, 62, 10, 63},
	block.YellowStainedGlassPane.ID:          []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.Lilac.ID:                           []ID{0},
	block.PurpleBanner.ID:                    []ID{0},
	block.MossyStoneBrickStairs.ID:           []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.StrippedCrimsonHyphae.ID:           []ID{1},
	block.SpruceSlab.ID:                      []ID{169, 169, 61, 61, 1, 1},
	block.GreenShulkerBox.ID:                 []ID{1},
	block.PolishedAndesiteSlab.ID:            []ID{169, 169, 61, 61, 1, 1},
	block.WeepingVinesPlant.ID:               []ID{0},
	block.CrackedNetherBricks.ID:             []ID{1},
	block.DiamondOre.ID:                      []ID{1},
	block.SkeletonSkull.ID:                   []ID{157},
	block.OrangeBanner.ID:                    []ID{0},
	block.LightBlueBanner.ID:                 []ID{0},
	block.AcaciaFenceGate.ID:                 []ID{0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75},
	block.BlackstoneWall.ID:                  []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.DeadFireCoralFan.ID:                []ID{0},
	block.DarkOakLeaves.ID:                   []ID{1},
	block.BrownWool.ID:                       []ID{1},
	block.BlueStainedGlassPane.ID:            []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.Terracotta.ID:                      []ID{1},
	block.SpruceDoor.ID:                      []ID{55, 55, 54, 54, 56, 56, 54, 54, 55, 55, 54, 54, 56, 56, 54, 54, 56, 56, 57, 57, 55, 55, 57, 57, 56, 56, 57, 57, 55, 55, 57, 57, 54, 54, 56, 56, 57, 57, 56, 56, 54, 54, 56, 56, 57, 57, 56, 56, 57, 57, 55, 55, 54, 54, 55, 55, 57, 57, 55, 55, 54, 54, 55, 55},
	block.PurpleShulkerBox.ID:                []ID{1},
	block.BlueConcretePowder.ID:              []ID{1},
	block.PottedCrimsonRoots.ID:              []ID{156},
	block.OrangeStainedGlassPane.ID:          []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.CoarseDirt.ID:                      []ID{1},
	block.Piston.ID:                          []ID{6, 7, 8, 9, 10, 11, 1, 1, 1, 1, 1, 1},
	block.SoulSand.ID:                        []ID{63},
	block.IronBars.ID:                        []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.DragonEgg.ID:                       []ID{111},
	block.BirchButton.ID:                     []ID{0},
	block.BrownTerracotta.ID:                 []ID{1},
	block.DarkOakSlab.ID:                     []ID{169, 169, 61, 61, 1, 1},
	block.TubeCoralFan.ID:                    []ID{0},
	block.WarpedDoor.ID:                      []ID{55, 55, 54, 54, 56, 56, 54, 54, 55, 55, 54, 54, 56, 56, 54, 54, 56, 56, 57, 57, 55, 55, 57, 57, 56, 56, 57, 57, 55, 55, 57, 57, 54, 54, 56, 56, 57, 57, 56, 56, 54, 54, 56, 56, 57, 57, 56, 56, 57, 57, 55, 55, 54, 54, 55, 55, 57, 57, 55, 55, 54, 54, 55, 55},
	block.PolishedBlackstone.ID:              []ID{1},
	block.Vine.ID:                            []ID{0},
	block.Cocoa.ID:                           []ID{112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123},
	block.CyanBanner.ID:                      []ID{0},
	block.Kelp.ID:                            []ID{0},
	block.Basalt.ID:                          []ID{1},
	block.WarpedTrapdoor.ID:                  []ID{54, 54, 54, 54, 89, 89, 89, 89, 54, 54, 54, 54, 88, 88, 88, 88, 57, 57, 57, 57, 89, 89, 89, 89, 57, 57, 57, 57, 88, 88, 88, 88, 56, 56, 56, 56, 89, 89, 89, 89, 56, 56, 56, 56, 88, 88, 88, 88, 55, 55, 55, 55, 89, 89, 89, 89, 55, 55, 55, 55, 88, 88, 88, 88},
	block.StrippedJungleLog.ID:               []ID{1},
	block.ActivatorRail.ID:                   []ID{0},
	block.BlackStainedGlassPane.ID:           []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.Prismarine.ID:                      []ID{1},
	block.RedstoneWallTorch.ID:               []ID{0},
	block.PottedWhiteTulip.ID:                []ID{156},
	block.LightGrayGlazedTerracotta.ID:       []ID{1},
	block.TurtleEgg.ID:                       []ID{238, 238, 238, 239, 239, 239, 239, 239, 239, 239, 239, 239},
	block.WarpedPressurePlate.ID:             []ID{0},
	block.Wheat.ID:                           []ID{0},
	block.PottedAllium.ID:                    []ID{156},
	block.PottedWitherRose.ID:                []ID{156},
	block.WhiteCarpet.ID:                     []ID{170},
	block.SweetBerryBush.ID:                  []ID{0},
	block.WarpedButton.ID:                    []ID{0},
	block.SmoothRedSandstone.ID:              []ID{1},
	block.LightGrayBed.ID:                    []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.GreenWool.ID:                       []ID{1},
	block.MushroomStem.ID:                    []ID{1},
	block.PottedPoppy.ID:                     []ID{156},
	block.PottedCornflower.ID:                []ID{156},
	block.PrismarineSlab.ID:                  []ID{169, 169, 61, 61, 1, 1},
	block.BlackWallBanner.ID:                 []ID{0},
	block.BoneBlock.ID:                       []ID{1},
	block.DeadHornCoralBlock.ID:              []ID{1},
	block.WetSponge.ID:                       []ID{1},
	block.LightBlueWool.ID:                   []ID{1},
	block.SprucePressurePlate.ID:             []ID{0},
	block.PurpleStainedGlass.ID:              []ID{1},
	block.DarkOakFence.ID:                    []ID{66, 67, 66, 67, 68, 69, 68, 69, 70, 71, 70, 71, 72, 73, 72, 73, 74, 75, 74, 75, 76, 77, 76, 77, 78, 79, 78, 79, 80, 65, 80, 65},
	block.MagmaBlock.ID:                      []ID{1},
	block.DeadFireCoralBlock.ID:              []ID{1},
	block.WarpedFungus.ID:                    []ID{0},
	block.AncientDebris.ID:                   []ID{1},
	block.BrainCoralFan.ID:                   []ID{0},
	block.GrassBlock.ID:                      []ID{1},
	block.Fire.ID:                            []ID{0},
	block.LilyPad.ID:                         []ID{106},
	block.RedTerracotta.ID:                   []ID{1},
	block.BlueCarpet.ID:                      []ID{170},
	block.StoneBrickSlab.ID:                  []ID{169, 169, 61, 61, 1, 1},
	block.HornCoralBlock.ID:                  []ID{1},
	block.RedNetherBrickStairs.ID:            []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.Barrel.ID:                          []ID{1},
	block.Lectern.ID:                         []ID{255},
	block.Stonecutter.ID:                     []ID{256},
	block.OakDoor.ID:                         []ID{55, 55, 54, 54, 56, 56, 54, 54, 55, 55, 54, 54, 56, 56, 54, 54, 56, 56, 57, 57, 55, 55, 57, 57, 56, 56, 57, 57, 55, 55, 57, 57, 54, 54, 56, 56, 57, 57, 56, 56, 54, 54, 56, 56, 57, 57, 56, 56, 57, 57, 55, 55, 54, 54, 55, 55, 57, 57, 55, 55, 54, 54, 55, 55},
	block.AttachedMelonStem.ID:               []ID{0},
	block.PottedDandelion.ID:                 []ID{156},
	block.LightGrayConcrete.ID:               []ID{1},
	block.WarpedHyphae.ID:                    []ID{1},
	block.OakSign.ID:                         []ID{0},
	block.AcaciaSign.ID:                      []ID{0},
	block.LimeStainedGlass.ID:                []ID{1},
	block.HayBlock.ID:                        []ID{1},
	block.WhiteConcrete.ID:                   []ID{1},
	block.Water.ID:                           []ID{0},
	block.PolishedGraniteStairs.ID:           []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.AndesiteWall.ID:                    []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.BlueWool.ID:                        []ID{1},
	block.SpruceButton.ID:                    []ID{0},
	block.LightWeightedPressurePlate.ID:      []ID{0},
	block.CyanShulkerBox.ID:                  []ID{1},
	block.YellowConcrete.ID:                  []ID{1},
	block.GildedBlackstone.ID:                []ID{1},
	block.LightGrayWallBanner.ID:             []ID{0},
	block.LimeShulkerBox.ID:                  []ID{1},
	block.DeadFireCoral.ID:                   []ID{0},
	block.PolishedAndesiteStairs.ID:          []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.WarpedPlanks.ID:                    []ID{1},
	block.OakWood.ID:                         []ID{1},
	block.WhiteWool.ID:                       []ID{1},
	block.BirchWallSign.ID:                   []ID{0},
	block.NetherBrickFence.ID:                []ID{66, 67, 66, 67, 68, 69, 68, 69, 70, 71, 70, 71, 72, 73, 72, 73, 74, 75, 74, 75, 76, 77, 76, 77, 78, 79, 78, 79, 80, 65, 80, 65},
	block.WitherSkeletonWallSkull.ID:         []ID{158, 159, 160, 161},
	block.BlueBanner.ID:                      []ID{0},
	block.StrippedDarkOakLog.ID:              []ID{1},
	block.Cobweb.ID:                          []ID{0},
	block.WhiteTulip.ID:                      []ID{0},
	block.DarkOakSign.ID:                     []ID{0},
	block.BlueTerracotta.ID:                  []ID{1},
	block.WhiteBanner.ID:                     []ID{0},
	block.LimeBanner.ID:                      []ID{0},
	block.BrownWallBanner.ID:                 []ID{0},
	block.VoidAir.ID:                         []ID{0},
	block.SoulSoil.ID:                        []ID{1},
	block.WarpedStem.ID:                      []ID{1},
	block.Obsidian.ID:                        []ID{1},
	block.BirchPressurePlate.ID:              []ID{0},
	block.Peony.ID:                           []ID{0},
	block.PinkConcretePowder.ID:              []ID{1},
	block.CrackedPolishedBlackstoneBricks.ID: []ID{1},
	block.CoalBlock.ID:                       []ID{1},
	block.TallGrass.ID:                       []ID{0},
	block.MagentaBed.ID:                      []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.PinkStainedGlassPane.ID:            []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.ChainCommandBlock.ID:               []ID{1},
	block.PolishedBlackstoneBricks.ID:        []ID{1},
	block.BirchLog.ID:                        []ID{1},
	block.SpruceFenceGate.ID:                 []ID{0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75},
	block.BirchFenceGate.ID:                  []ID{0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75},
	block.ShulkerBox.ID:                      []ID{1},
	block.Dandelion.ID:                       []ID{0},
	block.FireCoralWallFan.ID:                []ID{0},
	block.NetherSprouts.ID:                   []ID{0},
	block.CrimsonStem.ID:                     []ID{1},
	block.BirchFence.ID:                      []ID{66, 67, 66, 67, 68, 69, 68, 69, 70, 71, 70, 71, 72, 73, 72, 73, 74, 75, 74, 75, 76, 77, 76, 77, 78, 79, 78, 79, 80, 65, 80, 65},
	block.RedConcrete.ID:                     []ID{1},
	block.EndStoneBrickSlab.ID:               []ID{169, 169, 61, 61, 1, 1},
	block.DarkOakLog.ID:                      []ID{1},
	block.ChiseledSandstone.ID:               []ID{1},
	block.PottedBirchSapling.ID:              []ID{156},
	block.Anvil.ID:                           []ID{162, 162, 163, 163},
	block.MagentaShulkerBox.ID:               []ID{1},
	block.BlueGlazedTerracotta.ID:            []ID{1},
	block.BubbleColumn.ID:                    []ID{0},
	block.Lava.ID:                            []ID{0},
	block.Clay.ID:                            []ID{1},
	block.MossyStoneBricks.ID:                []ID{1},
	block.PinkBanner.ID:                      []ID{0},
	block.WhiteWallBanner.ID:                 []ID{0},
	block.BlackShulkerBox.ID:                 []ID{1},
	block.Chain.ID:                           []ID{270},
	block.Sandstone.ID:                       []ID{1},
	block.PinkCarpet.ID:                      []ID{170},
	block.CrimsonFence.ID:                    []ID{66, 67, 66, 67, 68, 69, 68, 69, 70, 71, 70, 71, 72, 73, 72, 73, 74, 75, 74, 75, 76, 77, 76, 77, 78, 79, 78, 79, 80, 65, 80, 65},
	block.Poppy.ID:                           []ID{0},
	block.CommandBlock.ID:                    []ID{1},
	block.CyanGlazedTerracotta.ID:            []ID{1},
	block.Cake.ID:                            []ID{81, 82, 83, 84, 85, 86, 87},
	block.OrangeTerracotta.ID:                []ID{1},
	block.GreenStainedGlassPane.ID:           []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.RedWallBanner.ID:                   []ID{0},
	block.ChorusFlower.ID:                    []ID{1},
	block.WhiteConcretePowder.ID:             []ID{1},
	block.CrimsonHyphae.ID:                   []ID{1},
	block.Lantern.ID:                         []ID{267, 266},
	block.Sand.ID:                            []ID{1},
	block.IronOre.ID:                         []ID{1},
	block.StonePressurePlate.ID:              []ID{0},
	block.RedMushroomBlock.ID:                []ID{1},
	block.JungleStairs.ID:                    []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.Observer.ID:                        []ID{1},
	block.PinkGlazedTerracotta.ID:            []ID{1},
	block.LightBlueConcretePowder.ID:         []ID{1},
	block.JungleLog.ID:                       []ID{1},
	block.OakFence.ID:                        []ID{66, 67, 66, 67, 68, 69, 68, 69, 70, 71, 70, 71, 72, 73, 72, 73, 74, 75, 74, 75, 76, 77, 76, 77, 78, 79, 78, 79, 80, 65, 80, 65},
	block.LightBlueStainedGlass.ID:           []ID{1},
	block.YellowTerracotta.ID:                []ID{1},
	block.RedSandstoneStairs.ID:              []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.SmoothStoneSlab.ID:                 []ID{169, 169, 61, 61, 1, 1},
	block.JungleDoor.ID:                      []ID{55, 55, 54, 54, 56, 56, 54, 54, 55, 55, 54, 54, 56, 56, 54, 54, 56, 56, 57, 57, 55, 55, 57, 57, 56, 56, 57, 57, 55, 55, 57, 57, 54, 54, 56, 56, 57, 57, 56, 56, 54, 54, 56, 56, 57, 57, 56, 56, 57, 57, 55, 55, 54, 54, 55, 55, 57, 57, 55, 55, 54, 54, 55, 55},
	block.DeadBubbleCoralFan.ID:              []ID{0},
	block.BlueIce.ID:                         []ID{1},
	block.Bell.ID:                            []ID{257, 257, 257, 257, 258, 258, 258, 258, 259, 259, 259, 259, 259, 259, 259, 259, 260, 260, 261, 261, 262, 262, 263, 263, 264, 264, 264, 264, 265, 265, 265, 265},
	block.CrimsonPressurePlate.ID:            []ID{0},
	block.PolishedGranite.ID:                 []ID{1},
	block.JunglePlanks.ID:                    []ID{1},
	block.BirchLeaves.ID:                     []ID{1},
	block.PottedBrownMushroom.ID:             []ID{156},
	block.PurpleWallBanner.ID:                []ID{0},
	block.LimeGlazedTerracotta.ID:            []ID{1},
	block.OakPlanks.ID:                       []ID{1},
	block.GoldBlock.ID:                       []ID{1},
	block.MossyCobblestone.ID:                []ID{1},
	block.WitherSkeletonSkull.ID:             []ID{157},
	block.YellowBanner.ID:                    []ID{0},
	block.YellowWallBanner.ID:                []ID{0},
	block.BrickSlab.ID:                       []ID{169, 169, 61, 61, 1, 1},
	block.Spawner.ID:                         []ID{1},
	block.PottedDarkOakSapling.ID:            []ID{156},
	block.ZombieHead.ID:                      []ID{157},
	block.DarkOakFenceGate.ID:                []ID{0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75},
	block.BlackGlazedTerracotta.ID:           []ID{1},
	block.AndesiteStairs.ID:                  []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.RespawnAnchor.ID:                   []ID{1},
	block.PurpleBed.ID:                       []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.InfestedCobblestone.ID:             []ID{1},
	block.SkeletonWallSkull.ID:               []ID{158, 159, 160, 161},
	block.ChippedAnvil.ID:                    []ID{162, 162, 163, 163},
	block.SmoothRedSandstoneStairs.ID:        []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.StoneStairs.ID:                     []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.BrownMushroom.ID:                   []ID{0},
	block.DarkOakTrapdoor.ID:                 []ID{54, 54, 54, 54, 89, 89, 89, 89, 54, 54, 54, 54, 88, 88, 88, 88, 57, 57, 57, 57, 89, 89, 89, 89, 57, 57, 57, 57, 88, 88, 88, 88, 56, 56, 56, 56, 89, 89, 89, 89, 56, 56, 56, 56, 88, 88, 88, 88, 55, 55, 55, 55, 89, 89, 89, 89, 55, 55, 55, 55, 88, 88, 88, 88},
	block.PlayerHead.ID:                      []ID{157},
	block.MagentaStainedGlassPane.ID:         []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.SlimeBlock.ID:                      []ID{1},
	block.JungleSlab.ID:                      []ID{169, 169, 61, 61, 1, 1},
	block.HornCoral.ID:                       []ID{0},
	block.NetherBricks.ID:                    []ID{1},
	block.RedstoneLamp.ID:                    []ID{1},
	block.PottedLilyOfTheValley.ID:           []ID{156},
	block.NetherQuartzOre.ID:                 []ID{1},
	block.LimeCarpet.ID:                      []ID{170},
	block.AcaciaDoor.ID:                      []ID{55, 55, 54, 54, 56, 56, 54, 54, 55, 55, 54, 54, 56, 56, 54, 54, 56, 56, 57, 57, 55, 55, 57, 57, 56, 56, 57, 57, 55, 55, 57, 57, 54, 54, 56, 56, 57, 57, 56, 56, 54, 54, 56, 56, 57, 57, 56, 56, 57, 57, 55, 55, 54, 54, 55, 55, 57, 57, 55, 55, 54, 54, 55, 55},
	block.WarpedWartBlock.ID:                 []ID{1},
	block.NoteBlock.ID:                       []ID{1},
	block.LilyOfTheValley.ID:                 []ID{0},
	block.Glowstone.ID:                       []ID{1},
	block.PolishedBlackstonePressurePlate.ID: []ID{0},
	block.GrayWool.ID:                        []ID{1},
	block.CrackedStoneBricks.ID:              []ID{1},
	block.PurpleCarpet.ID:                    []ID{170},
	block.WarpedFenceGate.ID:                 []ID{0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75},
	block.LapisBlock.ID:                      []ID{1},
	block.PistonHead.ID:                      []ID{13, 13, 12, 12, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20, 21, 21, 22, 22, 23, 23},
	block.MagentaStainedGlass.ID:             []ID{1},
	block.InfestedMossyStoneBricks.ID:        []ID{1},
	block.MagentaConcrete.ID:                 []ID{1},
	block.BubbleCoral.ID:                     []ID{0},
	block.InfestedCrackedStoneBricks.ID:      []ID{1},
	block.GlassPane.ID:                       []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.FlowerPot.ID:                       []ID{156},
	block.DamagedAnvil.ID:                    []ID{162, 162, 163, 163},
	block.CyanTerracotta.ID:                  []ID{1},
	block.DeadTubeCoralWallFan.ID:            []ID{0},
	block.LimeWool.ID:                        []ID{1},
	block.StoneBrickStairs.ID:                []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.BrownBanner.ID:                     []ID{0},
	block.GreenBanner.ID:                     []ID{0},
	block.RedSandstone.ID:                    []ID{1},
	block.PolishedBlackstoneStairs.ID:        []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.GoldOre.ID:                         []ID{1},
	block.CutSandstone.ID:                    []ID{1},
	block.Repeater.ID:                        []ID{58},
	block.StrippedCrimsonStem.ID:             []ID{1},
	block.CrimsonDoor.ID:                     []ID{55, 55, 54, 54, 56, 56, 54, 54, 55, 55, 54, 54, 56, 56, 54, 54, 56, 56, 57, 57, 55, 55, 57, 57, 56, 56, 57, 57, 55, 55, 57, 57, 54, 54, 56, 56, 57, 57, 56, 56, 54, 54, 56, 56, 57, 57, 56, 56, 57, 57, 55, 55, 54, 54, 55, 55, 57, 57, 55, 55, 54, 54, 55, 55},
	block.HoneyBlock.ID:                      []ID{64},
	block.SpruceWood.ID:                      []ID{1},
	block.WhiteBed.ID:                        []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.Chest.ID:                           []ID{48, 48, 49, 49, 50, 50, 48, 48, 50, 50, 49, 49, 48, 48, 51, 51, 52, 52, 48, 48, 52, 52, 51, 51},
	block.RedStainedGlass.ID:                 []ID{1},
	block.PottedAcaciaSapling.ID:             []ID{156},
	block.PottedRedMushroom.ID:               []ID{156},
	block.PurpleStainedGlassPane.ID:          []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.CryingObsidian.ID:                  []ID{1},
	block.Glass.ID:                           []ID{1},
	block.DetectorRail.ID:                    []ID{0},
	block.AcaciaPressurePlate.ID:             []ID{0},
	block.EndPortalFrame.ID:                  []ID{110, 110, 110, 110, 109, 109, 109, 109},
	block.PlayerWallHead.ID:                  []ID{158, 159, 160, 161},
	block.RedNetherBrickSlab.ID:              []ID{169, 169, 61, 61, 1, 1},
	block.GrayStainedGlass.ID:                []ID{1},
	block.AcaciaButton.ID:                    []ID{0},
	block.BirchSlab.ID:                       []ID{169, 169, 61, 61, 1, 1},
	block.AcaciaPlanks.ID:                    []ID{1},
	block.Seagrass.ID:                        []ID{0},
	block.NetherWart.ID:                      []ID{0},
	block.DeadFireCoralWallFan.ID:            []ID{0},
	block.BrownGlazedTerracotta.ID:           []ID{1},
	block.Bamboo.ID:                          []ID{245},
	block.MossyCobblestoneStairs.ID:          []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.RedSandstoneWall.ID:                []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.FletchingTable.ID:                  []ID{1},
	block.Blackstone.ID:                      []ID{1},
	block.BirchTrapdoor.ID:                   []ID{54, 54, 54, 54, 89, 89, 89, 89, 54, 54, 54, 54, 88, 88, 88, 88, 57, 57, 57, 57, 89, 89, 89, 89, 57, 57, 57, 57, 88, 88, 88, 88, 56, 56, 56, 56, 89, 89, 89, 89, 56, 56, 56, 56, 88, 88, 88, 88, 55, 55, 55, 55, 89, 89, 89, 89, 55, 55, 55, 55, 88, 88, 88, 88},
	block.PottedFern.ID:                      []ID{156},
	block.OakSlab.ID:                         []ID{169, 169, 61, 61, 1, 1},
	block.MossyStoneBrickSlab.ID:             []ID{169, 169, 61, 61, 1, 1},
	block.AndesiteSlab.ID:                    []ID{169, 169, 61, 61, 1, 1},
	block.CrimsonFenceGate.ID:                []ID{0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75},
	block.NetheriteBlock.ID:                  []ID{1},
	block.AttachedPumpkinStem.ID:             []ID{0},
	block.PottedJungleSapling.ID:             []ID{156},
	block.AcaciaStairs.ID:                    []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.RoseBush.ID:                        []ID{0},
	block.CutRedSandstone.ID:                 []ID{1},
	block.PurpleConcrete.ID:                  []ID{1},
	block.Composter.ID:                       []ID{269},
	block.Jukebox.ID:                         []ID{1},
	block.OrangeStainedGlass.ID:              []ID{1},
	block.BlackStainedGlass.ID:               []ID{1},
	block.ChiseledQuartzBlock.ID:             []ID{1},
	block.GrassPath.ID:                       []ID{53},
	block.SmoothRedSandstoneSlab.ID:          []ID{169, 169, 61, 61, 1, 1},
	block.BrickWall.ID:                       []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.GrayBanner.ID:                      []ID{0},
	block.AcaciaLog.ID:                       []ID{1},
	block.GreenBed.ID:                        []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.PoweredRail.ID:                     []ID{0},
	block.RedstoneWire.ID:                    []ID{0},
	block.JungleSign.ID:                      []ID{0},
	block.Cauldron.ID:                        []ID{108},
	block.Comparator.ID:                      []ID{58},
	block.StrippedDarkOakWood.ID:             []ID{1},
	block.CutRedSandstoneSlab.ID:             []ID{169, 169, 61, 61, 1, 1},
	block.NetherGoldOre.ID:                   []ID{1},
	block.RedTulip.ID:                        []ID{0},
	block.PottedSpruceSapling.ID:             []ID{156},
	block.CreeperWallHead.ID:                 []ID{158, 159, 160, 161},
	block.CyanCarpet.ID:                      []ID{170},
	block.SmoothSandstoneStairs.ID:           []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.BlackBanner.ID:                     []ID{0},
	block.StrippedOakWood.ID:                 []ID{1},
	block.RedBed.ID:                          []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.RedstoneOre.ID:                     []ID{1},
	block.SugarCane.ID:                       []ID{0},
	block.StoneBricks.ID:                     []ID{1},
	block.EnderChest.ID:                      []ID{48},
	block.CreeperHead.ID:                     []ID{157},
	block.NetherPortal.ID:                    []ID{0},
	block.YellowStainedGlass.ID:              []ID{1},
	block.PurpurSlab.ID:                      []ID{169, 169, 61, 61, 1, 1},
	block.EndRod.ID:                          []ID{172, 173, 172, 173, 171, 171},
	block.SmoothQuartzStairs.ID:              []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.StrippedWarpedStem.ID:              []ID{1},
	block.RepeatingCommandBlock.ID:           []ID{1},
	block.RedSand.ID:                         []ID{1},
	block.YellowWool.ID:                      []ID{1},
	block.PinkTulip.ID:                       []ID{0},
	block.EndPortal.ID:                       []ID{0},
	block.LightBlueStainedGlassPane.ID:       []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.GrayCarpet.ID:                      []ID{170},
	block.RedCarpet.ID:                       []ID{170},
	block.YellowGlazedTerracotta.ID:          []ID{1},
	block.DeadHornCoral.ID:                   []ID{0},
	block.ChiseledPolishedBlackstone.ID:      []ID{1},
	block.Dirt.ID:                            []ID{1},
	block.StrippedBirchLog.ID:                []ID{1},
	block.OrangeTulip.ID:                     []ID{0},
	block.MagentaCarpet.ID:                   []ID{170},
	block.AcaciaFence.ID:                     []ID{66, 67, 66, 67, 68, 69, 68, 69, 70, 71, 70, 71, 72, 73, 72, 73, 74, 75, 74, 75, 76, 77, 76, 77, 78, 79, 78, 79, 80, 65, 80, 65},
	block.GreenConcrete.ID:                   []ID{1},
	block.FireCoralFan.ID:                    []ID{0},
	block.ChiseledStoneBricks.ID:             []ID{1},
	block.SandstoneSlab.ID:                   []ID{169, 169, 61, 61, 1, 1},
	block.BirchDoor.ID:                       []ID{55, 55, 54, 54, 56, 56, 54, 54, 55, 55, 54, 54, 56, 56, 54, 54, 56, 56, 57, 57, 55, 55, 57, 57, 56, 56, 57, 57, 55, 55, 57, 57, 54, 54, 56, 56, 57, 57, 56, 56, 54, 54, 56, 56, 57, 57, 56, 56, 57, 57, 55, 55, 54, 54, 55, 55, 57, 57, 55, 55, 54, 54, 55, 55},
	block.LimeConcrete.ID:                    []ID{1},
	block.SeaPickle.ID:                       []ID{240, 240, 241, 241, 242, 242, 243, 243},
	block.Granite.ID:                         []ID{1},
	block.Rail.ID:                            []ID{0},
	block.EmeraldBlock.ID:                    []ID{1},
	block.DarkPrismarine.ID:                  []ID{1},
	block.FireCoralBlock.ID:                  []ID{1},
	block.SmoothQuartzSlab.ID:                []ID{169, 169, 61, 61, 1, 1},
	block.WallTorch.ID:                       []ID{0},
	block.WeepingVines.ID:                    []ID{0},
	block.OakWallSign.ID:                     []ID{0},
	block.Carrots.ID:                         []ID{0},
	block.PurpurBlock.ID:                     []ID{1},
	block.BlueConcrete.ID:                    []ID{1},
	block.GrayConcretePowder.ID:              []ID{1},
	block.PottedBamboo.ID:                    []ID{156},
	block.BlackWool.ID:                       []ID{1},
	block.QuartzPillar.ID:                    []ID{1},
	block.BlueWallBanner.ID:                  []ID{0},
	block.RedSandstoneSlab.ID:                []ID{169, 169, 61, 61, 1, 1},
	block.BrainCoral.ID:                      []ID{0},
	block.HornCoralFan.ID:                    []ID{0},
	block.CrimsonButton.ID:                   []ID{0},
	block.PottedDeadBush.ID:                  []ID{156},
	block.WhiteTerracotta.ID:                 []ID{1},
	block.PrismarineBrickSlab.ID:             []ID{169, 169, 61, 61, 1, 1},
	block.LightBlueCarpet.ID:                 []ID{170},
	block.CyanConcretePowder.ID:              []ID{1},
	block.Sponge.ID:                          []ID{1},
	block.EndStoneBricks.ID:                  []ID{1},
	block.LightGrayShulkerBox.ID:             []ID{1},
	block.PolishedBlackstoneBrickStairs.ID:   []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.Bricks.ID:                          []ID{1},
	block.JungleButton.ID:                    []ID{0},
	block.BambooSapling.ID:                   []ID{0},
	block.GraniteStairs.ID:                   []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.Campfire.ID:                        []ID{268},
	block.SoulLantern.ID:                     []ID{267, 266},
	block.RedMushroom.ID:                     []ID{0},
	block.LightGrayCarpet.ID:                 []ID{170},
	block.SoulTorch.ID:                       []ID{0},
	block.WarpedRoots.ID:                     []ID{0},
	block.CrimsonNylium.ID:                   []ID{1},
	block.DeadBrainCoralBlock.ID:             []ID{1},
	block.JungleWood.ID:                      []ID{1},
	block.OakPressurePlate.ID:                []ID{0},
	block.WhiteStainedGlass.ID:               []ID{1},
	block.GrayStainedGlassPane.ID:            []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.LightGrayStainedGlassPane.ID:       []ID{91, 92, 91, 92, 93, 94, 93, 94, 95, 96, 95, 96, 97, 98, 97, 98, 99, 100, 99, 100, 101, 102, 101, 102, 103, 104, 103, 104, 105, 90, 105, 90},
	block.AcaciaSlab.ID:                      []ID{169, 169, 61, 61, 1, 1},
	block.StructureVoid.ID:                   []ID{0},
	block.PolishedGraniteSlab.ID:             []ID{169, 169, 61, 61, 1, 1},
	block.StrippedWarpedHyphae.ID:            []ID{1},
	block.CrimsonPlanks.ID:                   []ID{1},
	block.Shroomlight.ID:                     []ID{1},
	block.MovingPiston.ID:                    []ID{0},
	block.CyanStainedGlass.ID:                []ID{1},
	block.BrownStainedGlass.ID:               []ID{1},
	block.OakFenceGate.ID:                    []ID{0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 72, 72, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75, 0, 0, 75, 75},
	block.PottedOxeyeDaisy.ID:                []ID{156},
	block.RedNetherBrickWall.ID:              []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.SoulCampfire.ID:                    []ID{268},
	block.StickyPiston.ID:                    []ID{6, 7, 8, 9, 10, 11, 1, 1, 1, 1, 1, 1},
	block.GreenStainedGlass.ID:               []ID{1},
	block.DaylightDetector.ID:                []ID{60},
	block.RedConcretePowder.ID:               []ID{1},
	block.BrainCoralBlock.ID:                 []ID{1},
	block.BlackstoneSlab.ID:                  []ID{1},
	block.RedstoneTorch.ID:                   []ID{0},
	block.EndStoneBrickWall.ID:               []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.RedShulkerBox.ID:                   []ID{1},
	block.Andesite.ID:                        []ID{1},
	block.PolishedAndesite.ID:                []ID{1},
	block.SpruceLeaves.ID:                    []ID{1},
	block.Potatoes.ID:                        []ID{0},
	block.PrismarineBricks.ID:                []ID{1},
	block.SpruceFence.ID:                     []ID{66, 67, 66, 67, 68, 69, 68, 69, 70, 71, 70, 71, 72, 73, 72, 73, 74, 75, 74, 75, 76, 77, 76, 77, 78, 79, 78, 79, 80, 65, 80, 65},
	block.BlueShulkerBox.ID:                  []ID{1},
	block.GrayGlazedTerracotta.ID:            []ID{1},
	block.BlueStainedGlass.ID:                []ID{1},
	block.MelonStem.ID:                       []ID{0},
	block.PottedBlueOrchid.ID:                []ID{156},
	block.CrimsonRoots.ID:                    []ID{0},
	block.BirchSapling.ID:                    []ID{0},
	block.DragonWallHead.ID:                  []ID{158, 159, 160, 161},
	block.TubeCoral.ID:                       []ID{0},
	block.DioriteWall.ID:                     []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.DarkOakWood.ID:                     []ID{1},
	block.BlackBed.ID:                        []ID{3, 2, 3, 2, 2, 3, 2, 3, 4, 5, 4, 5, 5, 4, 5, 4},
	block.OakStairs.ID:                       []ID{25, 25, 26, 26, 27, 27, 28, 28, 29, 29, 24, 24, 30, 30, 31, 31, 32, 32, 33, 33, 34, 34, 35, 35, 36, 36, 37, 37, 38, 38, 39, 39, 40, 40, 41, 41, 42, 42, 43, 43, 44, 44, 36, 36, 26, 26, 38, 38, 28, 28, 45, 45, 41, 41, 30, 30, 43, 43, 32, 32, 46, 46, 27, 27, 35, 35, 29, 29, 37, 37, 47, 47, 31, 31, 40, 40, 33, 33, 42, 42},
	block.PinkConcrete.ID:                    []ID{1},
	block.BrownConcretePowder.ID:             []ID{1},
	block.DeadBubbleCoralBlock.ID:            []ID{1},
	block.PottedCrimsonFungus.ID:             []ID{156},
	block.PottedOrangeTulip.ID:               []ID{156},
	block.BlackCarpet.ID:                     []ID{170},
	block.Lodestone.ID:                       []ID{1},
	block.PolishedBlackstoneWall.ID:          []ID{125, 126, 125, 126, 127, 128, 127, 128, 129, 130, 129, 130, 131, 132, 131, 132, 133, 134, 133, 134, 135, 136, 135, 136, 137, 138, 137, 138, 139, 140, 139, 140, 141, 142, 141, 142, 143, 144, 143, 144, 145, 146, 145, 146, 147, 148, 147, 148, 149, 150, 149, 150, 151, 152, 151, 152, 153, 124, 153, 124, 154, 155, 154, 155},
	block.JungleTrapdoor.ID:                  []ID{54, 54, 54, 54, 89, 89, 89, 89, 54, 54, 54, 54, 88, 88, 88, 88, 57, 57, 57, 57, 89, 89, 89, 89, 57, 57, 57, 57, 88, 88, 88, 88, 56, 56, 56, 56, 89, 89, 89, 89, 56, 56, 56, 56, 88, 88, 88, 88, 55, 55, 55, 55, 89, 89, 89, 89, 55, 55, 55, 55, 88, 88, 88, 88},
	block.Sunflower.ID:                       []ID{0},
	block.LimeWallBanner.ID:                  []ID{0},
	block.WhiteGlazedTerracotta.ID:           []ID{1},
}

// Dimensions describes the bounding boxes of a shape ID.
var Dimensions = map[ID]Shape{
	122: Shape{
		ID: 122,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0.1875, Z: 0.25},
				Max: BoundingTriplet{X: 0.5625, Y: 0.75, Z: 0.75},
			},
		},
	},
	6: Shape{
		ID: 6,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	37: Shape{
		ID: 37,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
		},
	},
	39: Shape{
		ID: 39,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	105: Shape{
		ID: 105,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.5625},
			},
		},
	},
	27: Shape{
		ID: 27,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
		},
	},
	82: Shape{
		ID: 82,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.5, Z: 0.9375},
			},
		},
	},
	258: Shape{
		ID: 258,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.75, Y: 1, Z: 1},
			},
		},
	},
	64: Shape{
		ID: 64,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.9375, Z: 0.9375},
			},
		},
	},
	246: Shape{
		ID: 246,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.125, Y: 1, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.875},
				Max: BoundingTriplet{X: 0.125, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.875, Z: 0.125},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.875, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.875, Z: 0.875},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.875, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.875, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0.875, Y: 0, Z: 0.875},
				Max: BoundingTriplet{X: 1, Y: 0.875, Z: 1},
			},
		},
	},
	15: Shape{
		ID: 15,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: -0.25, Y: 0.375, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 0.625, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.375, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.375},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0.625},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0.625, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.625},
			},
		},
	},
	145: Shape{
		ID: 145,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.25},
			},
		},
	},
	215: Shape{
		ID: 215,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	269: Shape{
		ID: 269,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.125, Z: 0},
				Max: BoundingTriplet{X: 0.125, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.125, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.125, Z: 0.875},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.875, Y: 0.125, Z: 0.125},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.875},
			},
		},
	},
	83: Shape{
		ID: 83,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.5, Z: 0.9375},
			},
		},
	},
	113: Shape{
		ID: 113,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.4375, Z: 0.6875},
				Max: BoundingTriplet{X: 0.625, Y: 0.75, Z: 0.9375},
			},
		},
	},
	209: Shape{
		ID: 209,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	226: Shape{
		ID: 226,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	33: Shape{
		ID: 33,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
		},
	},
	95: Shape{
		ID: 95,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0.5625},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 1},
			},
		},
	},
	125: Shape{
		ID: 125,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	146: Shape{
		ID: 146,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.25},
			},
		},
	},
	262: Shape{
		ID: 262,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.8125, Z: 0.4375},
				Max: BoundingTriplet{X: 0.8125, Y: 0.9375, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.375, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.8125, Z: 0.6875},
			},
		},
	},
	270: Shape{
		ID: 270,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.40625, Y: 0, Z: 0.40625},
				Max: BoundingTriplet{X: 0.59375, Y: 1, Z: 0.59375},
			},
		},
	},
	19: Shape{
		ID: 19,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.25, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.375, Z: 0.375},
				Max: BoundingTriplet{X: 1.25, Y: 0.625, Z: 0.625},
			},
		},
	},
	140: Shape{
		ID: 140,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
		},
	},
	184: Shape{
		ID: 184,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	244: Shape{
		ID: 244,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.3125, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.6875, Z: 0.6875},
			},
		},
	},
	199: Shape{
		ID: 199,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	237: Shape{
		ID: 237,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	238: Shape{
		ID: 238,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.75, Y: 0.4375, Z: 0.75},
			},
		},
	},
	154: Shape{
		ID: 154,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.6875},
			},
		},
	},
	170: Shape{
		ID: 170,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.0625, Z: 1},
			},
		},
	},
	229: Shape{
		ID: 229,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
		},
	},
	51: Shape{
		ID: 51,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.9375, Y: 0.875, Z: 0.9375},
			},
		},
	},
	66: Shape{
		ID: 66,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 0.375},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.625},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 1},
			},
		},
	},
	118: Shape{
		ID: 118,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0.3125, Z: 0.3125},
				Max: BoundingTriplet{X: 0.4375, Y: 0.75, Z: 0.6875},
			},
		},
	},
	148: Shape{
		ID: 148,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.6875},
			},
		},
	},
	61: Shape{
		ID: 61,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
		},
	},
	81: Shape{
		ID: 81,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.5, Z: 0.9375},
			},
		},
	},
	101: Shape{
		ID: 101,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.4375},
			},
		},
	},
	202: Shape{
		ID: 202,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	17: Shape{
		ID: 17,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.375, Z: -0.25},
				Max: BoundingTriplet{X: 0.625, Y: 0.625, Z: 0.75},
			},
		},
	},
	25: Shape{
		ID: 25,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	29: Shape{
		ID: 29,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 0.5},
			},
		},
	},
	49: Shape{
		ID: 49,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 1, Y: 0.875, Z: 0.9375},
			},
		},
	},
	12: Shape{
		ID: 12,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.375, Z: 0.25},
				Max: BoundingTriplet{X: 0.625, Y: 0.625, Z: 1.25},
			},
		},
	},
	124: Shape{
		ID: 124,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
		},
	},
	99: Shape{
		ID: 99,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.4375},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0.5625},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 1},
			},
		},
	},
	189: Shape{
		ID: 189,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
		},
	},
	248: Shape{
		ID: 248,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.125, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.25, Y: 0.8125, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.4375, Z: 0.3125},
				Max: BoundingTriplet{X: 0.875, Y: 0.8125, Z: 0.375},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.4375, Z: 0.625},
				Max: BoundingTriplet{X: 0.875, Y: 0.8125, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.125},
				Max: BoundingTriplet{X: 0.75, Y: 0.4375, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.4375, Z: 0.125},
				Max: BoundingTriplet{X: 0.75, Y: 1, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.4375, Z: 0.375},
				Max: BoundingTriplet{X: 0.875, Y: 0.8125, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.4375, Z: 0.6875},
				Max: BoundingTriplet{X: 0.75, Y: 1, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.8125, Z: 0.3125},
				Max: BoundingTriplet{X: 0.75, Y: 1, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.875, Y: 0.4375, Z: 0.625},
			},
		},
	},
	5: Shape{
		ID: 5,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5625, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.1875, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0, Z: 0.8125},
				Max: BoundingTriplet{X: 1, Y: 0.1875, Z: 1},
			},
		},
	},
	97: Shape{
		ID: 97,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5625},
			},
		},
	},
	230: Shape{
		ID: 230,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	234: Shape{
		ID: 234,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	10: Shape{
		ID: 10,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.75, Z: 1},
			},
		},
	},
	67: Shape{
		ID: 67,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.625, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.625},
			},
		},
	},
	96: Shape{
		ID: 96,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5625, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5625},
			},
		},
	},
	250: Shape{
		ID: 250,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.3125, Z: 0.4375},
				Max: BoundingTriplet{X: 0.875, Y: 0.6875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.375, Z: 0},
				Max: BoundingTriplet{X: 0.25, Y: 0.625, Z: 0.4375},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.125, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.3125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.3125, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.875, Z: 0.4375},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.3125, Z: 0.8125},
				Max: BoundingTriplet{X: 0.75, Y: 0.875, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.6875, Z: 0.4375},
				Max: BoundingTriplet{X: 0.75, Y: 0.875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0},
				Max: BoundingTriplet{X: 0.875, Y: 0.625, Z: 0.4375},
			},
		},
	},
	200: Shape{
		ID: 200,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	232: Shape{
		ID: 232,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	84: Shape{
		ID: 84,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.5, Z: 0.9375},
			},
		},
	},
	108: Shape{
		ID: 108,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.125, Y: 1, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 0.125, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.25},
				Max: BoundingTriplet{X: 1, Y: 0.25, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.125, Y: 1, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.25, Y: 1, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0, Z: 0.875},
				Max: BoundingTriplet{X: 0.25, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.1875, Z: 0.125},
				Max: BoundingTriplet{X: 1, Y: 0.25, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.1875, Z: 0.75},
				Max: BoundingTriplet{X: 1, Y: 0.25, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.1875, Z: 0.875},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.1875, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0, Z: 0.875},
				Max: BoundingTriplet{X: 1, Y: 0.1875, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.875, Y: 0, Z: 0.125},
				Max: BoundingTriplet{X: 1, Y: 0.1875, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.875, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 1, Y: 0.1875, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.875, Y: 0.25, Z: 0.125},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.875},
			},
		},
	},
	121: Shape{
		ID: 121,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.1875, Z: 0.4375},
				Max: BoundingTriplet{X: 0.75, Y: 0.75, Z: 0.9375},
			},
		},
	},
	195: Shape{
		ID: 195,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	90: Shape{
		ID: 90,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.5625},
			},
		},
	},
	251: Shape{
		ID: 251,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.125, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.875, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.3125, Z: 0.125},
				Max: BoundingTriplet{X: 0.5625, Y: 0.6875, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.3125, Z: 0.75},
				Max: BoundingTriplet{X: 0.5625, Y: 0.6875, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.5625, Y: 0.375, Z: 0.125},
				Max: BoundingTriplet{X: 1, Y: 0.625, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.5625, Y: 0.375, Z: 0.75},
				Max: BoundingTriplet{X: 1, Y: 0.625, Z: 0.875},
			},
		},
	},
	183: Shape{
		ID: 183,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	204: Shape{
		ID: 204,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	268: Shape{
		ID: 268,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.4375, Z: 1},
			},
		},
	},
	106: Shape{
		ID: 106,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.09375, Z: 0.9375},
			},
		},
	},
	233: Shape{
		ID: 233,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	235: Shape{
		ID: 235,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	147: Shape{
		ID: 147,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.3125},
			},
		},
	},
	156: Shape{
		ID: 156,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.375, Z: 0.6875},
			},
		},
	},
	171: Shape{
		ID: 171,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 1, Z: 0.625},
			},
		},
	},
	214: Shape{
		ID: 214,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	175: Shape{
		ID: 175,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	240: Shape{
		ID: 240,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 0.375, Z: 0.625},
			},
		},
	},
	35: Shape{
		ID: 35,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 0.5},
			},
		},
	},
	54: Shape{
		ID: 54,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.8125},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	88: Shape{
		ID: 88,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.1875, Z: 1},
			},
		},
	},
	151: Shape{
		ID: 151,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	193: Shape{
		ID: 193,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	228: Shape{
		ID: 228,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	261: Shape{
		ID: 261,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.375, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.8125, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.5625, Y: 0.9375, Z: 1},
			},
		},
	},
	69: Shape{
		ID: 69,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.625, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.625},
			},
		},
	},
	93: Shape{
		ID: 93,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.4375},
			},
		},
	},
	178: Shape{
		ID: 178,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	181: Shape{
		ID: 181,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
		},
	},
	87: Shape{
		ID: 87,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.5, Z: 0.9375},
			},
		},
	},
	132: Shape{
		ID: 132,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.6875, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
		},
	},
	155: Shape{
		ID: 155,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.5, Y: 0, Z: 0.5},
				Max: BoundingTriplet{X: 0.5, Y: 1.5, Z: 0.5},
			},
		},
	},
	201: Shape{
		ID: 201,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	163: Shape{
		ID: 163,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.625, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0, Z: 0.125},
				Max: BoundingTriplet{X: 0.875, Y: 0.25, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.8125, Y: 0.3125, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.3125, Z: 0.375},
				Max: BoundingTriplet{X: 0.75, Y: 0.625, Z: 0.625},
			},
		},
	},
	203: Shape{
		ID: 203,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	213: Shape{
		ID: 213,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
		},
	},
	63: Shape{
		ID: 63,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.875, Z: 1},
			},
		},
	},
	91: Shape{
		ID: 91,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.4375},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0.5625},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 1},
			},
		},
	},
	126: Shape{
		ID: 126,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
		},
	},
	211: Shape{
		ID: 211,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	60: Shape{
		ID: 60,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.375, Z: 1},
			},
		},
	},
	161: Shape{
		ID: 161,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.5, Y: 0.75, Z: 0.75},
			},
		},
	},
	58: Shape{
		ID: 58,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.125, Z: 1},
			},
		},
	},
	86: Shape{
		ID: 86,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.6875, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.5, Z: 0.9375},
			},
		},
	},
	120: Shape{
		ID: 120,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.1875, Z: 0.0625},
				Max: BoundingTriplet{X: 0.75, Y: 0.75, Z: 0.5625},
			},
		},
	},
	185: Shape{
		ID: 185,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	9: Shape{
		ID: 9,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	172: Shape{
		ID: 172,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.375, Z: 0},
				Max: BoundingTriplet{X: 0.625, Y: 0.625, Z: 1},
			},
		},
	},
	249: Shape{
		ID: 249,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.875, Y: 1, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.4375, Z: 0.125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.8125, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.4375, Z: 0.75},
				Max: BoundingTriplet{X: 0.6875, Y: 0.8125, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.125},
				Max: BoundingTriplet{X: 0.625, Y: 0.4375, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 0.625, Y: 0.4375, Z: 0.875},
			},
		},
	},
	266: Shape{
		ID: 266,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.4375, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.4375, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 0.5625, Z: 0.625},
			},
		},
	},
	45: Shape{
		ID: 45,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 0.5, Y: 1, Z: 1},
			},
		},
	},
	117: Shape{
		ID: 117,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.3125, Z: 0.5625},
				Max: BoundingTriplet{X: 0.6875, Y: 0.75, Z: 0.9375},
			},
		},
	},
	217: Shape{
		ID: 217,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	265: Shape{
		ID: 265,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.8125, Z: 0.4375},
				Max: BoundingTriplet{X: 1, Y: 0.9375, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.375, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.8125, Z: 0.6875},
			},
		},
	},
	260: Shape{
		ID: 260,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.375, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.8125, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0.8125, Z: 0},
				Max: BoundingTriplet{X: 0.5625, Y: 0.9375, Z: 0.8125},
			},
		},
	},
	263: Shape{
		ID: 263,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.4375},
				Max: BoundingTriplet{X: 1, Y: 0.9375, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.375, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.8125, Z: 0.6875},
			},
		},
	},
	62: Shape{
		ID: 62,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.625, Z: 1},
			},
		},
	},
	111: Shape{
		ID: 111,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 1, Z: 0.9375},
			},
		},
	},
	127: Shape{
		ID: 127,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	192: Shape{
		ID: 192,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	227: Shape{
		ID: 227,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	256: Shape{
		ID: 256,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5625, Z: 1},
			},
		},
	},
	55: Shape{
		ID: 55,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.1875, Y: 1, Z: 1},
			},
		},
	},
	107: Shape{
		ID: 107,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.125, Z: 0.9375},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0.125, Z: 0.4375},
				Max: BoundingTriplet{X: 0.5625, Y: 0.875, Z: 0.5625},
			},
		},
	},
	169: Shape{
		ID: 169,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	174: Shape{
		ID: 174,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	7: Shape{
		ID: 7,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.75, Y: 1, Z: 1},
			},
		},
	},
	11: Shape{
		ID: 11,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.25, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	141: Shape{
		ID: 141,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	257: Shape{
		ID: 257,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.75},
			},
		},
	},
	44: Shape{
		ID: 44,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	177: Shape{
		ID: 177,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	223: Shape{
		ID: 223,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	208: Shape{
		ID: 208,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	252: Shape{
		ID: 252,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.375, Z: 0.125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.625, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.375, Z: 0.75},
				Max: BoundingTriplet{X: 0.8125, Y: 0.625, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.125, Z: 0.25},
				Max: BoundingTriplet{X: 1, Y: 0.875, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0.3125, Z: 0.125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.375, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0.3125, Z: 0.75},
				Max: BoundingTriplet{X: 0.8125, Y: 0.375, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0.625, Z: 0.125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.6875, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0.625, Z: 0.75},
				Max: BoundingTriplet{X: 0.8125, Y: 0.6875, Z: 0.875},
			},
		},
	},
	8: Shape{
		ID: 8,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.75},
			},
		},
	},
	18: Shape{
		ID: 18,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.25, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.375, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 0.625, Z: 0.625},
			},
		},
	},
	21: Shape{
		ID: 21,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.75, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: -0.25, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 0.75, Z: 0.625},
			},
		},
	},
	160: Shape{
		ID: 160,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 1, Y: 0.75, Z: 0.75},
			},
		},
	},
	264: Shape{
		ID: 264,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.375, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.8125, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0.8125, Z: 0},
				Max: BoundingTriplet{X: 0.5625, Y: 0.9375, Z: 1},
			},
		},
	},
	14: Shape{
		ID: 14,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.375, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 0.625, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.375, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.375},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0.625},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0.625, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.625},
			},
		},
	},
	70: Shape{
		ID: 70,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.625},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 1},
			},
		},
	},
	142: Shape{
		ID: 142,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	167: Shape{
		ID: 167,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.25, Z: 0.375},
				Max: BoundingTriplet{X: 0.75, Y: 0.5, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.625, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.6875, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.6875, Z: 0},
				Max: BoundingTriplet{X: 0.125, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.6875, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.6875, Z: 0.875},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.625, Z: 0.375},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.625},
				Max: BoundingTriplet{X: 0.75, Y: 0.625, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.5, Z: 0.375},
				Max: BoundingTriplet{X: 0.75, Y: 0.625, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.875, Y: 0.6875, Z: 0.125},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.875},
			},
		},
	},
	20: Shape{
		ID: 20,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.75, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 0.75, Z: 0.625},
			},
		},
	},
	53: Shape{
		ID: 53,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.9375, Z: 1},
			},
		},
	},
	92: Shape{
		ID: 92,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5625, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5625},
			},
		},
	},
	139: Shape{
		ID: 139,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
		},
	},
	98: Shape{
		ID: 98,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5625},
			},
		},
	},
	187: Shape{
		ID: 187,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	1: Shape{
		ID: 1,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	3: Shape{
		ID: 3,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.1875, Y: 0.5625, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.5625, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5625, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.1875, Z: 0.1875},
			},
		},
	},
	74: Shape{
		ID: 74,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 0.375},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.625},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 1},
			},
		},
	},
	76: Shape{
		ID: 76,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 0.375},
			},
		},
	},
	2: Shape{
		ID: 2,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.8125},
				Max: BoundingTriplet{X: 0.1875, Y: 0.5625, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5625, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 1, Y: 0.5625, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0, Z: 0.8125},
				Max: BoundingTriplet{X: 1, Y: 0.1875, Z: 1},
			},
		},
	},
	24: Shape{
		ID: 24,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
		},
	},
	43: Shape{
		ID: 43,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0.5},
				Max: BoundingTriplet{X: 0.5, Y: 1, Z: 1},
			},
		},
	},
	162: Shape{
		ID: 162,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.125, Y: 0, Z: 0.125},
				Max: BoundingTriplet{X: 0.875, Y: 0.25, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.625, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.1875},
				Max: BoundingTriplet{X: 0.75, Y: 0.3125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.3125, Z: 0.25},
				Max: BoundingTriplet{X: 0.625, Y: 0.625, Z: 0.75},
			},
		},
	},
	100: Shape{
		ID: 100,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 1},
			},
		},
	},
	164: Shape{
		ID: 164,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.625, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.6875, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.6875, Z: 0},
				Max: BoundingTriplet{X: 0.125, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.6875, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.6875, Z: 0.875},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.625, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 0.25, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.875, Y: 0.6875, Z: 0.125},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.875},
			},
		},
	},
	242: Shape{
		ID: 242,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.125, Y: 0, Z: 0.125},
				Max: BoundingTriplet{X: 0.875, Y: 0.375, Z: 0.875},
			},
		},
	},
	4: Shape{
		ID: 4,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.1875, Y: 0.5625, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.8125},
				Max: BoundingTriplet{X: 0.1875, Y: 0.5625, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.5625, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5625, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 1, Y: 0.5625, Z: 1},
			},
		},
	},
	26: Shape{
		ID: 26,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	57: Shape{
		ID: 57,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.1875},
			},
		},
	},
	68: Shape{
		ID: 68,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 0.375},
			},
		},
	},
	0: Shape{
		ID:    0,
		Boxes: []BoundingBox{},
	},
	50: Shape{
		ID: 50,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.875, Z: 0.9375},
			},
		},
	},
	165: Shape{
		ID: 165,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.625, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.6875, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.6875, Z: 0},
				Max: BoundingTriplet{X: 0.125, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.6875, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.6875, Z: 0.875},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.625, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.25, Z: 0},
				Max: BoundingTriplet{X: 0.625, Y: 0.5, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.875, Y: 0.6875, Z: 0.125},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.875},
			},
		},
	},
	191: Shape{
		ID: 191,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	159: Shape{
		ID: 159,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0},
				Max: BoundingTriplet{X: 0.75, Y: 0.75, Z: 0.5},
			},
		},
	},
	219: Shape{
		ID: 219,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	243: Shape{
		ID: 243,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.125, Y: 0, Z: 0.125},
				Max: BoundingTriplet{X: 0.875, Y: 0.4375, Z: 0.875},
			},
		},
	},
	73: Shape{
		ID: 73,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.625},
			},
		},
	},
	104: Shape{
		ID: 104,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 1},
			},
		},
	},
	119: Shape{
		ID: 119,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.5625, Y: 0.3125, Z: 0.3125},
				Max: BoundingTriplet{X: 0.9375, Y: 0.75, Z: 0.6875},
			},
		},
	},
	152: Shape{
		ID: 152,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	166: Shape{
		ID: 166,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.625, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.6875, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.6875, Z: 0},
				Max: BoundingTriplet{X: 0.125, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.6875, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.6875, Z: 0.875},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.625, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.25, Z: 0.75},
				Max: BoundingTriplet{X: 0.625, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.875, Y: 0.6875, Z: 0.125},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.875},
			},
		},
	},
	46: Shape{
		ID: 46,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
		},
	},
	52: Shape{
		ID: 52,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.875, Z: 1},
			},
		},
	},
	80: Shape{
		ID: 80,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 0.625},
			},
		},
	},
	130: Shape{
		ID: 130,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
		},
	},
	42: Shape{
		ID: 42,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	253: Shape{
		ID: 253,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.1875, Z: 0.3125},
				Max: BoundingTriplet{X: 0.875, Y: 0.5625, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.5625, Z: 0.375},
				Max: BoundingTriplet{X: 0.875, Y: 0.75, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.75, Z: 0.375},
				Max: BoundingTriplet{X: 0.25, Y: 1, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.125},
				Max: BoundingTriplet{X: 0.75, Y: 0.1875, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.1875, Z: 0.125},
				Max: BoundingTriplet{X: 0.75, Y: 0.75, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.1875, Z: 0.6875},
				Max: BoundingTriplet{X: 0.75, Y: 0.75, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.5625, Z: 0.3125},
				Max: BoundingTriplet{X: 0.75, Y: 0.75, Z: 0.375},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.5625, Z: 0.625},
				Max: BoundingTriplet{X: 0.75, Y: 0.75, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0.75, Z: 0.375},
				Max: BoundingTriplet{X: 0.875, Y: 1, Z: 0.625},
			},
		},
	},
	206: Shape{
		ID: 206,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	225: Shape{
		ID: 225,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	131: Shape{
		ID: 131,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.3125},
			},
		},
	},
	133: Shape{
		ID: 133,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	173: Shape{
		ID: 173,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.375, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 0.625, Z: 0.625},
			},
		},
	},
	198: Shape{
		ID: 198,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
		},
	},
	176: Shape{
		ID: 176,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	182: Shape{
		ID: 182,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	231: Shape{
		ID: 231,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	254: Shape{
		ID: 254,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.125, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.875, Y: 0.75, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.1875, Z: 0.125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.5625, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.1875, Z: 0.75},
				Max: BoundingTriplet{X: 0.6875, Y: 0.5625, Z: 0.875},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.5625, Z: 0.125},
				Max: BoundingTriplet{X: 0.625, Y: 1, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.5625, Z: 0.75},
				Max: BoundingTriplet{X: 0.625, Y: 1, Z: 0.875},
			},
		},
	},
	34: Shape{
		ID: 34,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
		},
	},
	114: Shape{
		ID: 114,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0.4375, Z: 0.375},
				Max: BoundingTriplet{X: 0.3125, Y: 0.75, Z: 0.625},
			},
		},
	},
	138: Shape{
		ID: 138,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
		},
	},
	168: Shape{
		ID: 168,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.625, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.6875, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.6875, Z: 0},
				Max: BoundingTriplet{X: 0.125, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.6875, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.125},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.6875, Z: 0.875},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.625, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0.25, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.875, Y: 0.6875, Z: 0.125},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.875},
			},
		},
	},
	112: Shape{
		ID: 112,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.4375, Z: 0.0625},
				Max: BoundingTriplet{X: 0.625, Y: 0.75, Z: 0.3125},
			},
		},
	},
	129: Shape{
		ID: 129,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.25},
			},
		},
	},
	150: Shape{
		ID: 150,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	212: Shape{
		ID: 212,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	16: Shape{
		ID: 16,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.375, Z: 0},
				Max: BoundingTriplet{X: 0.625, Y: 0.625, Z: 0.75},
			},
		},
	},
	28: Shape{
		ID: 28,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5, Y: 1, Z: 0.5},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
		},
	},
	85: Shape{
		ID: 85,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.5625, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.5, Z: 0.9375},
			},
		},
	},
	110: Shape{
		ID: 110,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.8125, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1, Z: 0.75},
			},
		},
	},
	267: Shape{
		ID: 267,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.0625, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.5, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 0.625, Z: 0.625},
			},
		},
	},
	197: Shape{
		ID: 197,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
		},
	},
	23: Shape{
		ID: 23,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.25, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.25, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 1.25, Z: 0.625},
			},
		},
	},
	36: Shape{
		ID: 36,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
		},
	},
	157: Shape{
		ID: 157,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.5, Z: 0.75},
			},
		},
	},
	186: Shape{
		ID: 186,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	205: Shape{
		ID: 205,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
		},
	},
	40: Shape{
		ID: 40,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
		},
	},
	59: Shape{
		ID: 59,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.25, Z: 1},
			},
		},
	},
	103: Shape{
		ID: 103,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0.5625},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 1},
			},
		},
	},
	128: Shape{
		ID: 128,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.6875, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
		},
	},
	116: Shape{
		ID: 116,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.3125, Z: 0.0625},
				Max: BoundingTriplet{X: 0.6875, Y: 0.75, Z: 0.4375},
			},
		},
	},
	179: Shape{
		ID: 179,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	32: Shape{
		ID: 32,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 0.5, Y: 1, Z: 0.5},
			},
		},
	},
	136: Shape{
		ID: 136,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.6875, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
		},
	},
	180: Shape{
		ID: 180,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	222: Shape{
		ID: 222,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	48: Shape{
		ID: 48,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.875, Z: 0.9375},
			},
		},
	},
	102: Shape{
		ID: 102,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.5625},
			},
		},
	},
	137: Shape{
		ID: 137,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
		},
	},
	224: Shape{
		ID: 224,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	22: Shape{
		ID: 22,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.25, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.25, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 1, Z: 0.625},
			},
		},
	},
	221: Shape{
		ID: 221,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	220: Shape{
		ID: 220,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	38: Shape{
		ID: 38,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.5},
				Max: BoundingTriplet{X: 0.5, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	94: Shape{
		ID: 94,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.5625, Y: 0, Z: 0.4375},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5625},
			},
		},
	},
	135: Shape{
		ID: 135,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	144: Shape{
		ID: 144,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	31: Shape{
		ID: 31,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	79: Shape{
		ID: 79,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 1},
			},
		},
	},
	196: Shape{
		ID: 196,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
		},
	},
	241: Shape{
		ID: 241,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.375, Z: 0.8125},
			},
		},
	},
	149: Shape{
		ID: 149,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	247: Shape{
		ID: 247,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.3125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.875, Y: 0.6875, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.125, Y: 0.375, Z: 0.5625},
				Max: BoundingTriplet{X: 0.25, Y: 0.625, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.125, Z: 0},
				Max: BoundingTriplet{X: 0.75, Y: 0.3125, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.3125, Z: 0},
				Max: BoundingTriplet{X: 0.75, Y: 0.875, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.3125, Z: 0.5625},
				Max: BoundingTriplet{X: 0.75, Y: 0.875, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.6875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.75, Y: 0.875, Z: 0.5625},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0.5625},
				Max: BoundingTriplet{X: 0.875, Y: 0.625, Z: 1},
			},
		},
	},
	65: Shape{
		ID: 65,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 0.625},
			},
		},
	},
	71: Shape{
		ID: 71,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.625, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.625},
			},
		},
	},
	78: Shape{
		ID: 78,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 0.625},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0.625},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 1},
			},
		},
	},
	89: Shape{
		ID: 89,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.8125, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	190: Shape{
		ID: 190,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	207: Shape{
		ID: 207,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	218: Shape{
		ID: 218,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	239: Shape{
		ID: 239,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.0625, Y: 0, Z: 0.0625},
				Max: BoundingTriplet{X: 0.9375, Y: 0.4375, Z: 0.9375},
			},
		},
	},
	41: Shape{
		ID: 41,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 0.5, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0.5},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	72: Shape{
		ID: 72,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.375},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.625},
			},
		},
	},
	143: Shape{
		ID: 143,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
		},
	},
	158: Shape{
		ID: 158,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.5},
				Max: BoundingTriplet{X: 0.75, Y: 0.75, Z: 1},
			},
		},
	},
	109: Shape{
		ID: 109,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 1},
			},
		},
	},
	194: Shape{
		ID: 194,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 0.1875},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.8125},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
		},
	},
	210: Shape{
		ID: 210,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	255: Shape{
		ID: 255,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.125, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.875, Z: 0.75},
			},
		},
	},
	13: Shape{
		ID: 13,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.25},
			},
			{
				Min: BoundingTriplet{X: 0.375, Y: 0.375, Z: 0.25},
				Max: BoundingTriplet{X: 0.625, Y: 0.625, Z: 1},
			},
		},
	},
	115: Shape{
		ID: 115,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.6875, Y: 0.4375, Z: 0.375},
				Max: BoundingTriplet{X: 0.9375, Y: 0.75, Z: 0.625},
			},
		},
	},
	123: Shape{
		ID: 123,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0.1875, Z: 0.25},
				Max: BoundingTriplet{X: 0.9375, Y: 0.75, Z: 0.75},
			},
		},
	},
	236: Shape{
		ID: 236,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
		},
	},
	47: Shape{
		ID: 47,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	134: Shape{
		ID: 134,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0, Z: 0.75},
				Max: BoundingTriplet{X: 0.6875, Y: 1.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.75, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 1, Y: 1.5, Z: 0.6875},
			},
		},
	},
	245: Shape{
		ID: 245,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.15625, Y: 0, Z: 0.15625},
				Max: BoundingTriplet{X: 0.34375, Y: 1, Z: 0.34375},
			},
		},
	},
	259: Shape{
		ID: 259,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.25, Y: 0.25, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 0.375, Z: 0.75},
			},
			{
				Min: BoundingTriplet{X: 0.3125, Y: 0.375, Z: 0.3125},
				Max: BoundingTriplet{X: 0.6875, Y: 0.8125, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.4375, Y: 0.8125, Z: 0.4375},
				Max: BoundingTriplet{X: 0.5625, Y: 1, Z: 0.5625},
			},
		},
	},
	56: Shape{
		ID: 56,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 1},
			},
		},
	},
	75: Shape{
		ID: 75,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 1},
			},
		},
	},
	153: Shape{
		ID: 153,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0.3125},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.6875},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.25},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.3125},
			},
			{
				Min: BoundingTriplet{X: 0.25, Y: 0, Z: 0.6875},
				Max: BoundingTriplet{X: 0.75, Y: 1.5, Z: 0.75},
			},
		},
	},
	216: Shape{
		ID: 216,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 0.8125, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0.8125, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	30: Shape{
		ID: 30,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 0.5, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 0.5, Y: 1, Z: 1},
			},
			{
				Min: BoundingTriplet{X: 0.5, Y: 0.5, Z: 0},
				Max: BoundingTriplet{X: 1, Y: 1, Z: 0.5},
			},
		},
	},
	188: Shape{
		ID: 188,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.1875, Y: 0, Z: 0.1875},
				Max: BoundingTriplet{X: 0.8125, Y: 1, Z: 0.8125},
			},
			{
				Min: BoundingTriplet{X: 0.8125, Y: 0.1875, Z: 0.1875},
				Max: BoundingTriplet{X: 1, Y: 0.8125, Z: 0.8125},
			},
		},
	},
	77: Shape{
		ID: 77,
		Boxes: []BoundingBox{
			{
				Min: BoundingTriplet{X: 0.375, Y: 0, Z: 0},
				Max: BoundingTriplet{X: 0.625, Y: 1.5, Z: 0.625},
			},
		},
	},
}
