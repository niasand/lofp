package engine

import (
	"fmt"
	"strings"

	"github.com/jonradoff/lofp/i18n"
)

// emoteEntry defines the self/room messages for an emote.
// Placeholders: %N = actor first name, %P = his/her, %O = him/her, %E = he/she
type emoteEntry struct {
	Self       string // what the actor sees (no target)
	Room       string // what the room sees (no target)
	SelfTarget string // what the actor sees (with target) — %T = target name
	RoomTarget string // what the room sees (with target) — %T = target name
}

var emoteTable = map[string]emoteEntry{
	"SMILE":     {Self: i18n.T("You smile."), Room: i18n.T("%N smiles."), SelfTarget: i18n.T("You smile at %T."), RoomTarget: i18n.T("%N smiles at %T.")},
	"BOW":       {Self: i18n.T("You bow."), Room: i18n.T("%N bows."), SelfTarget: i18n.T("You bow to %T."), RoomTarget: i18n.T("%N bows to %T.")},
	"CURTSEY":   {Self: i18n.T("You curtsey."), Room: i18n.T("%N curtseys."), SelfTarget: i18n.T("You curtsey to %T."), RoomTarget: i18n.T("%N curtseys to %T.")},
	"CURTSY":    {Self: i18n.T("You curtsy."), Room: i18n.T("%N curtsies."), SelfTarget: i18n.T("You curtsy to %T."), RoomTarget: i18n.T("%N curtsies to %T.")},
	"WAVE":      {Self: i18n.T("You wave."), Room: i18n.T("%N waves."), SelfTarget: i18n.T("You wave to %T."), RoomTarget: i18n.T("%N waves to %T.")},
	"NOD":       {Self: i18n.T("You nod."), Room: i18n.T("%N nods."), SelfTarget: i18n.T("You nod to %T."), RoomTarget: i18n.T("%N nods to %T.")},
	"LAUGH":     {Self: i18n.T("You laugh."), Room: i18n.T("%N laughs."), SelfTarget: i18n.T("You laugh at %T."), RoomTarget: i18n.T("%N laughs at %T.")},
	"CHUCKLE":   {Self: i18n.T("You chuckle."), Room: i18n.T("%N chuckles."), SelfTarget: i18n.T("You chuckle at %T."), RoomTarget: i18n.T("%N chuckles at %T.")},
	"GRIN":      {Self: i18n.T("You grin."), Room: i18n.T("%N grins."), SelfTarget: i18n.T("You grin at %T."), RoomTarget: i18n.T("%N grins at %T.")},
	"FROWN":     {Self: i18n.T("You frown."), Room: i18n.T("%N frowns."), SelfTarget: i18n.T("You frown at %T."), RoomTarget: i18n.T("%N frowns at %T.")},
	"SIGH":      {Self: i18n.T("You sigh."), Room: i18n.T("%N sighs."), SelfTarget: i18n.T("You sigh at %T."), RoomTarget: i18n.T("%N sighs at %T.")},
	"SHRUG":     {Self: i18n.T("You shrug."), Room: i18n.T("%N shrugs."), SelfTarget: i18n.T("You shrug at %T."), RoomTarget: i18n.T("%N shrugs at %T.")},
	"WINK":      {Self: i18n.T("You wink."), Room: i18n.T("%N winks."), SelfTarget: i18n.T("You wink at %T."), RoomTarget: i18n.T("%N winks at %T.")},
	"CRY":       {Self: i18n.T("You cry."), Room: i18n.T("%N cries."), SelfTarget: i18n.T("You cry on %T's shoulder."), RoomTarget: i18n.T("%N cries on %T's shoulder.")},
	"DANCE":     {Self: i18n.T("You dance."), Room: i18n.T("%N dances."), SelfTarget: i18n.T("You dance with %T."), RoomTarget: i18n.T("%N dances with %T.")},
	"HUG":       {Self: i18n.T("You hug yourself."), Room: i18n.T("%N hugs themselves."), SelfTarget: i18n.T("You hug %T."), RoomTarget: i18n.T("%N hugs %T.")},
	"KISS":      {Self: i18n.T("You blow a kiss."), Room: i18n.T("%N blows a kiss."), SelfTarget: i18n.T("You kiss %T."), RoomTarget: i18n.T("%N kisses %T.")},
	"POKE":      {Self: i18n.T("You poke yourself."), Room: i18n.T("%N pokes themselves."), SelfTarget: i18n.T("You poke %T."), RoomTarget: i18n.T("%N pokes %T.")},
	"TICKLE":    {Self: i18n.T("You tickle yourself."), Room: i18n.T("%N tickles themselves."), SelfTarget: i18n.T("You tickle %T."), RoomTarget: i18n.T("%N tickles %T.")},
	"SLAP":      {Self: i18n.T("You slap yourself."), Room: i18n.T("%N slaps themselves."), SelfTarget: i18n.T("You slap %T."), RoomTarget: i18n.T("%N slaps %T.")},
	"HOWL":      {Self: i18n.T("You howl."), Room: i18n.T("%N howls."), SelfTarget: i18n.T("You howl at %T."), RoomTarget: i18n.T("%N howls at %T.")},
	"SING":      {Self: i18n.T("You sing."), Room: i18n.T("%N sings."), SelfTarget: i18n.T("You sing to %T."), RoomTarget: i18n.T("%N sings to %T.")},
	"PACE":      {Self: i18n.T("You pace back and forth."), Room: i18n.T("%N paces back and forth.")},
	"FIDGET":    {Self: i18n.T("You fidget."), Room: i18n.T("%N fidgets.")},
	"SHIVER":    {Self: i18n.T("You shiver."), Room: i18n.T("%N shivers.")},
	"SNORT":     {Self: i18n.T("You snort."), Room: i18n.T("%N snorts."), SelfTarget: i18n.T("You snort at %T."), RoomTarget: i18n.T("%N snorts at %T.")},
	"GROAN":     {Self: i18n.T("You groan."), Room: i18n.T("%N groans.")},
	"MUMBLE":    {Self: i18n.T("You mumble something."), Room: i18n.T("%N mumbles something.")},
	"BABBLE":    {Self: i18n.T("You babble."), Room: i18n.T("%N babbles.")},
	"BEAM":      {Self: i18n.T("You beam."), Room: i18n.T("%N beams."), SelfTarget: i18n.T("You beam at %T."), RoomTarget: i18n.T("%N beams at %T.")},
	"SWOON":     {Self: i18n.T("You swoon."), Room: i18n.T("%N swoons.")},
	"TOAST":     {Self: i18n.T("You raise your glass in a toast."), Room: i18n.T("%N raises a toast."), SelfTarget: i18n.T("You raise a toast to %T."), RoomTarget: i18n.T("%N raises a toast to %T.")},
	"SHUDDER":   {Self: i18n.T("You shudder."), Room: i18n.T("%N shudders.")},
	"POINT":     {Self: i18n.T("You point."), Room: i18n.T("%N points."), SelfTarget: i18n.T("You point at %T."), RoomTarget: i18n.T("%N points at %T.")},
	"KICK":      {Self: i18n.T("You kick at the ground."), Room: i18n.T("%N kicks at the ground."), SelfTarget: i18n.T("You kick %T."), RoomTarget: i18n.T("%N kicks %T.")},
	"KNOCK":     {Self: i18n.T("You knock."), Room: i18n.T("%N knocks."), SelfTarget: i18n.T("You knock on %T."), RoomTarget: i18n.T("%N knocks on %T.")},
	"TOUCH":     {Self: i18n.T("You touch yourself."), Room: i18n.T("%N touches themselves."), SelfTarget: i18n.T("You touch %T."), RoomTarget: i18n.T("%N touches %T.")},
	"RUB":       {Self: i18n.T("You rub your hands together."), Room: i18n.T("%N rubs %P hands together."), SelfTarget: i18n.T("You rub %T."), RoomTarget: i18n.T("%N rubs %T.")},
	"PET":       {Self: i18n.T("You pet yourself."), Room: i18n.T("%N pets themselves."), SelfTarget: i18n.T("You pet %T."), RoomTarget: i18n.T("%N pets %T.")},
	"PUNCH":     {Self: i18n.T("You punch the air."), Room: i18n.T("%N punches the air."), SelfTarget: i18n.T("You punch %T."), RoomTarget: i18n.T("%N punches %T.")},
	"SPIT":      {Self: i18n.T("You spit."), Room: i18n.T("%N spits."), SelfTarget: i18n.T("You spit at %T."), RoomTarget: i18n.T("%N spits at %T.")},
	"GAZE":      {Self: i18n.T("You gaze about."), Room: i18n.T("%N gazes about."), SelfTarget: i18n.T("You gaze at %T."), RoomTarget: i18n.T("%N gazes at %T.")},
	"GLARE":     {Self: i18n.T("You glare."), Room: i18n.T("%N glares."), SelfTarget: i18n.T("You glare at %T."), RoomTarget: i18n.T("%N glares at %T.")},
	"SCOWL":     {Self: i18n.T("You scowl."), Room: i18n.T("%N scowls."), SelfTarget: i18n.T("You scowl at %T."), RoomTarget: i18n.T("%N scowls at %T.")},
	"COMFORT":   {Self: i18n.T("You comfort yourself."), Room: i18n.T("%N comforts themselves."), SelfTarget: i18n.T("You comfort %T."), RoomTarget: i18n.T("%N comforts %T.")},
	"RECITE":    {Self: i18n.T("You recite."), Room: i18n.T("%N recites.")},
	"YAWN":      {Self: i18n.T("You yawn."), Room: i18n.T("%N yawns."), SelfTarget: i18n.T("You yawn at %T."), RoomTarget: i18n.T("%N yawns at %T.")},
	"BLINK":     {Self: i18n.T("You blink."), Room: i18n.T("%N blinks."), SelfTarget: i18n.T("You blink at %T."), RoomTarget: i18n.T("%N blinks at %T.")},
	"BLUSH":     {Self: i18n.T("You blush."), Room: i18n.T("%N blushes.")},
	"CRINGE":    {Self: i18n.T("You cringe."), Room: i18n.T("%N cringes.")},
	"CUDDLE":    {Self: i18n.T("You cuddle up."), Room: i18n.T("%N cuddles up."), SelfTarget: i18n.T("You cuddle up to %T."), RoomTarget: i18n.T("%N cuddles up to %T.")},
	"COUGH":     {Self: i18n.T("You cough."), Room: i18n.T("%N coughs.")},
	"FURROW":    {Self: i18n.T("You furrow your brow."), Room: i18n.T("%N furrows %P brow.")},
	"GASP":      {Self: i18n.T("You gasp."), Room: i18n.T("%N gasps.")},
	"GIGGLE":    {Self: i18n.T("You giggle."), Room: i18n.T("%N giggles."), SelfTarget: i18n.T("You giggle at %T."), RoomTarget: i18n.T("%N giggles at %T.")},
	"GRIMACE":   {Self: i18n.T("You grimace."), Room: i18n.T("%N grimaces.")},
	"GROWL":     {Self: i18n.T("You growl."), Room: i18n.T("%N growls."), SelfTarget: i18n.T("You growl at %T."), RoomTarget: i18n.T("%N growls at %T.")},
	"GULP":      {Self: i18n.T("You gulp."), Room: i18n.T("%N gulps.")},
	"JUMP":      {Self: i18n.T("You jump up and down."), Room: i18n.T("%N jumps up and down.")},
	"LEAN":      {Self: i18n.T("You lean back."), Room: i18n.T("%N leans back."), SelfTarget: i18n.T("You lean on %T."), RoomTarget: i18n.T("%N leans on %T.")},
	"NUZZLE":    {Self: i18n.T("You nuzzle."), Room: i18n.T("%N nuzzles."), SelfTarget: i18n.T("You nuzzle %T affectionately."), RoomTarget: i18n.T("%N nuzzles %T affectionately.")},
	"PANT":      {Self: i18n.T("You pant."), Room: i18n.T("%N pants.")},
	"PONDER":    {Self: i18n.T("You ponder for a moment."), Room: i18n.T("%N ponders.")},
	"POUT":      {Self: i18n.T("You pout."), Room: i18n.T("%N pouts.")},
	"ROLL":      {Self: i18n.T("You roll your eyes."), Room: i18n.T("%N rolls %P eyes."), SelfTarget: i18n.T("You roll your eyes at %T."), RoomTarget: i18n.T("%N rolls %P eyes at %T.")},
	"SCREAM":    {Self: i18n.T("You scream!"), Room: i18n.T("%N screams!")},
	"SMIRK":     {Self: i18n.T("You smirk."), Room: i18n.T("%N smirks."), SelfTarget: i18n.T("You smirk at %T."), RoomTarget: i18n.T("%N smirks at %T.")},
	"SNICKER":   {Self: i18n.T("You snicker."), Room: i18n.T("%N snickers.")},
	"SALUTE":    {Self: i18n.T("You salute."), Room: i18n.T("%N salutes."), SelfTarget: i18n.T("You salute %T."), RoomTarget: i18n.T("%N salutes %T.")},
	"STRETCH":   {Self: i18n.T("You stretch your arms lazily."), Room: i18n.T("%N stretches %P arms lazily.")},
	"TAP":       {Self: i18n.T("You tap your foot."), Room: i18n.T("%N taps %P foot."), SelfTarget: i18n.T("You tap %T on the shoulder."), RoomTarget: i18n.T("%N taps %T on the shoulder.")},
	"TWIRL":     {Self: i18n.T("You twirl around."), Room: i18n.T("%N twirls around.")},
	"WINCE":     {Self: i18n.T("You wince."), Room: i18n.T("%N winces.")},
	"WHISTLE":   {Self: i18n.T("You whistle innocently."), Room: i18n.T("%N whistles innocently.")},
	"MUTTER":    {Self: i18n.T("You mutter something under your breath."), Room: i18n.T("%N mutters something you can't quite make out.")},
	"CARESS":    {Self: i18n.T("You caress yourself."), Room: i18n.T("%N caresses themselves."), SelfTarget: i18n.T("You caress %T."), RoomTarget: i18n.T("%N caresses %T.")},
	"NUDGE":     {Self: i18n.T("You nudge."), Room: i18n.T("%N nudges."), SelfTarget: i18n.T("You nudge %T."), RoomTarget: i18n.T("%N nudges %T.")},
	"ARCH":      {Self: i18n.T("You arch an eyebrow."), Room: i18n.T("%N arches %P eyebrow."), SelfTarget: i18n.T("You arch an eyebrow at %T."), RoomTarget: i18n.T("%N arches %P eyebrow at %T.")},
	"RAISE":     {Self: i18n.T("You raise an eyebrow."), Room: i18n.T("%N raises an eyebrow."), SelfTarget: i18n.T("You raise an eyebrow towards %T."), RoomTarget: i18n.T("%N raises an eyebrow towards %T.")},
	"HEAD":      {Self: i18n.T("You shake your head."), Room: i18n.T("%N shakes %P head back and forth."), SelfTarget: i18n.T("You shake your head at %T."), RoomTarget: i18n.T("%N shakes %P head at %T.")},
	"SCRATCH":   {Self: i18n.T("You scratch your head."), Room: i18n.T("%N scratches %P head.")},
	"CLAP":      {Self: i18n.T("You clap."), Room: i18n.T("%N claps."), SelfTarget: i18n.T("You clap for %T."), RoomTarget: i18n.T("%N claps for %T.")},
	"SNIFF":     {Self: i18n.T("You sniff."), Room: i18n.T("%N sniffs."), SelfTarget: i18n.T("You sniff %T."), RoomTarget: i18n.T("%N sniffs %T.")},
	"LISTEN":    {Self: i18n.T("You listen carefully."), Room: i18n.T("%N listens carefully."), SelfTarget: i18n.T("You listen carefully to %T."), RoomTarget: i18n.T("%N listens carefully to %T.")},

	// Additional emotes from alphabetical list
	"LICK":      {Self: i18n.T("You lick your lips."), Room: i18n.T("%N licks %P lips."), SelfTarget: i18n.T("You kiss %T passionately."), RoomTarget: i18n.T("%N kisses %T passionately.")},
	"NIBBLE":    {Self: i18n.T("You nibble."), Room: i18n.T("%N nibbles."), SelfTarget: i18n.T("You nibble on %T."), RoomTarget: i18n.T("%N nibbles on %T.")},
	"BARK":      {Self: i18n.T("You bark."), Room: i18n.T("%N barks.")},
	"CLAW":      {Self: i18n.T("You claw at the air."), Room: i18n.T("%N claws at the air."), SelfTarget: i18n.T("You claw at %T."), RoomTarget: i18n.T("%N claws at %T.")},
	"CURSE":     {Self: i18n.T("You curse."), Room: i18n.T("%N curses.")},
	"DUCK":      {Self: i18n.T("You duck."), Room: i18n.T("%N ducks.")},
	"HISS":      {Self: i18n.T("You hiss."), Room: i18n.T("%N hisses."), SelfTarget: i18n.T("You hiss at %T."), RoomTarget: i18n.T("%N hisses at %T.")},
	"HOLD":      {Self: i18n.T("You hold your hands together."), Room: i18n.T("%N holds %P hands together."), SelfTarget: i18n.T("You hold %T."), RoomTarget: i18n.T("%N holds %T.")},
	"HULA":      {Self: i18n.T("You do a little hula dance."), Room: i18n.T("%N does a little hula dance.")},
	"JIG":       {Self: i18n.T("You do a little jig."), Room: i18n.T("%N does a little jig.")},
	"MOAN":      {Self: i18n.T("You moan."), Room: i18n.T("%N moans.")},
	"MASSAGE":   {Self: i18n.T("You massage your temples."), Room: i18n.T("%N massages %P temples."), SelfTarget: i18n.T("You massage %T."), RoomTarget: i18n.T("%N massages %T.")},
	"PINCH":     {Self: i18n.T("You pinch yourself."), Room: i18n.T("%N pinches themselves."), SelfTarget: i18n.T("You pinch %T."), RoomTarget: i18n.T("%N pinches %T.")},
	"PLAY":      {Self: i18n.T("You play around."), Room: i18n.T("%N plays around.")},
	"PURR":      {Self: i18n.T("You purr."), Room: i18n.T("%N purrs."), SelfTarget: i18n.T("You purr at %T."), RoomTarget: i18n.T("%N purrs at %T.")},
	"ROAR":      {Self: i18n.T("You roar!"), Room: i18n.T("%N roars!")},
	"SNARL":     {Self: i18n.T("You snarl."), Room: i18n.T("%N snarls."), SelfTarget: i18n.T("You snarl at %T."), RoomTarget: i18n.T("%N snarls at %T.")},
	"SNUGGLE":   {Self: i18n.T("You snuggle up."), Room: i18n.T("%N snuggles up."), SelfTarget: i18n.T("You snuggle up to %T."), RoomTarget: i18n.T("%N snuggles up to %T.")},
	"WAG":       {Self: i18n.T("You wag your finger."), Room: i18n.T("%N wags %P finger."), SelfTarget: i18n.T("You wag your finger at %T."), RoomTarget: i18n.T("%N wags %P finger at %T.")},
	"WAIT":      {Self: i18n.T("You wait patiently."), Room: i18n.T("%N waits patiently.")},
	"WRITE":     {Self: i18n.T("You write something."), Room: i18n.T("%N writes something.")},
	"YOWL":      {Self: i18n.T("You yowl!"), Room: i18n.T("%N yowls!")},
	"THUMP":     {Self: i18n.T("You thump the ground."), Room: i18n.T("%N thumps the ground."), SelfTarget: i18n.T("You thump %T."), RoomTarget: i18n.T("%N thumps %T.")},
	"STOMP":     {Self: i18n.T("You stomp your feet."), Room: i18n.T("%N stomps %p feet."), SelfTarget: i18n.T("You stomp towards %T."), RoomTarget: i18n.T("%N stomps towards %T.")},
	"APPLAUD":   {Self: i18n.T("You applaud."), Room: i18n.T("%N applauds."), SelfTarget: i18n.T("You applaud %T."), RoomTarget: i18n.T("%N applauds %T.")},
	"PEER":      {Self: i18n.T("You peer around."), Room: i18n.T("%N peers around."), SelfTarget: i18n.T("You peer at %T."), RoomTarget: i18n.T("%N peers at %T.")},
	"GRUNT":     {Self: i18n.T("You grunt."), Room: i18n.T("%N grunts.")},
	"DIP":       {Self: i18n.T("You dip."), Room: i18n.T("%N dips."), SelfTarget: i18n.T("You dip %T."), RoomTarget: i18n.T("%N dips %T.")},
	"HANDRAISE": {Self: i18n.T("You raise your hand."), Room: i18n.T("%N raises %P hand.")},
	"HANDSHAKE": {Self: i18n.T("You extend your hand."), Room: i18n.T("%N extends %P hand."), SelfTarget: i18n.T("You shake hands with %T."), RoomTarget: i18n.T("%N shakes hands with %T.")},
	"HEADSHAKE": {Self: i18n.T("You shake your head."), Room: i18n.T("%N shakes %P head.")},
	"PICK":      {Self: i18n.T("You pick at something."), Room: i18n.T("%N picks at something.")},
	"GESTURE":   {Self: i18n.T("You gesture."), Room: i18n.T("%N gestures."), SelfTarget: i18n.T("You gesture at %T."), RoomTarget: i18n.T("%N gestures at %T.")},
}

// selfOverrides maps verb+target to special self-targeted messages.
// Key format: "VERB:me" for self-targeting overrides.
var selfOverrides = map[string][2]string{
	"SPIT:me":   {i18n.T("You drool all over yourself."), i18n.T("%N drools all over themselves.")},
	"LICK:me":   {i18n.T("You lick your lips."), i18n.T("%N licks %P lips.")},
	"LAUGH:me":  {i18n.T("You laugh at yourself."), i18n.T("%N laughs at themselves.")},
	"KICK:me":   {i18n.T("You kick yourself."), i18n.T("%N kicks themselves.")},
	"THUMP:me":  {i18n.T("You thump yourself on the head."), i18n.T("%N thumps themselves on the head.")},
}

// raceEmotes are emotes only available to specific races. Key = "RACE:VERB".
var raceEmotes = map[string]emoteEntry{
	// Drakin
	"6:FLICK":   {Self: i18n.T("You flick your forked tongue."), Room: i18n.T("%N flicks %P forked tongue.")},
	"6:BARE":    {Self: i18n.T("You bare your teeth menacingly."), Room: i18n.T("%N bares %P teeth menacingly.")},
	"6:SPREAD":  {Self: i18n.T("You spread your wings wide."), Room: i18n.T("%N spreads %P wings wide.")},
	"6:FOLD":    {Self: i18n.T("You fold your wings against your body."), Room: i18n.T("%N folds %P wings.")},
	"6:SWISH":   {Self: i18n.T("You swish your tail."), Room: i18n.T("%N swishes %P tail.")},
	// Aelfen
	"2:RUBEARS": {Self: i18n.T("You rub the points of your ears."), Room: i18n.T("%N rubs the points of %P ears.")},
	// Highlander
	"3:PULLBEARD": {Self: i18n.T("You pull on your beard."), Room: i18n.T("%N pulls on %P beard.")},
	// Wolf form emotes (wolfling)
	"4:SCRATCH": {Self: i18n.T("You scratch behind your ear."), Room: i18n.T("%N scratches behind %P ear.")},
	"4:BARE":    {Self: i18n.T("You bare your fangs."), Room: i18n.T("%N bares %P fangs.")},
	"4:CHASE":   {Self: i18n.T("You chase your tail."), Room: i18n.T("%N chases %P tail.")},
	"4:SCENT":   {Self: i18n.T("You lift your nose and scent the air."), Room: i18n.T("%N lifts %P nose and scents the air.")},
	"4:WHINE":   {Self: i18n.T("You whine softly."), Room: i18n.T("%N whines softly.")},
	"4:DROOP":   {Self: i18n.T("You droop your tail."), Room: i18n.T("%N droops %P tail.")},
}

// additionalSelfEmotes — emotes with no target that aren't in the main table.
// These are triggered by specific verbs with no arguments.
var additionalSelfEmotes = map[string][2]string{
	"FUME":     {i18n.T("You fume."), i18n.T("%N fumes.")},
	"SQUINT":   {i18n.T("You squint your eyes."), i18n.T("%N squints %P eyes.")},
	"HUM":      {i18n.T("You hum softly to yourself."), i18n.T("%N hums softly.")},
	"SNIFFLE":  {i18n.T("You sniffle."), i18n.T("%N sniffles.")},
	"SLOUCH":   {i18n.T("You slouch."), i18n.T("%N slouches.")},
	"SNORE":    {i18n.T("You snore loudly."), i18n.T("%N snores loudly.")},
	"SNEEZE":   {i18n.T("You sneeze."), i18n.T("%N sneezes.")},
	"STARE":    {i18n.T("You stare off into space."), i18n.T("%N stares off into space.")},
	"PUCKER":   {i18n.T("You pucker your lips."), i18n.T("%N puckers %P lips.")},
	"CRACK":    {i18n.T("You crack your knuckles."), i18n.T("%N cracks %P knuckles.")},
	"BITE":     {i18n.T("You bite your lower lip."), i18n.T("%N bites %P lower lip.")},
	"BOUNCE":   {i18n.T("You bounce up and down."), i18n.T("%N bounces up and down.")},
	"STRIKE":   {i18n.T("You strike a heroic pose."), i18n.T("%N strikes a heroic pose.")},
	"CLUTCH":   {i18n.T("You clutch your head."), i18n.T("%N clutches %P head.")},
	"WIPE":     {i18n.T("You wipe your brow."), i18n.T("%N wipes %P brow.")},
	"GRIT":     {i18n.T("You grit your teeth."), i18n.T("%N grits %P teeth.")},
	"TOSS":     {i18n.T("You toss your hands up in the air."), i18n.T("%N tosses %P hands up in the air.")},
	"ATTENTION": {i18n.T("You stand at attention."), i18n.T("%N stands at attention.")},
	"TONGUE":   {i18n.T("You stick your tongue out."), i18n.T("%N sticks %P tongue out.")},
	"WRINKLE":  {i18n.T("You wrinkle your nose."), i18n.T("%N wrinkles %P nose.")},
	"PUFF":     {i18n.T("You puff out your cheeks."), i18n.T("%N puffs out %P cheeks.")},
	"DIZZY":    {i18n.T("You feel dizzy."), i18n.T("%N looks dizzy.")},
	"BAT":      {i18n.T("You bat your eyelashes."), i18n.T("%N bats %P eyelashes.")},
}

// kissBodyParts defines the body parts that can be kissed and whether they require submit.
var kissBodyParts = map[string]struct {
	requiresSubmit bool
	selfMsg        string
	roomMsg        string
	targetMsg      string
}{
	"head":  {false, i18n.T("You kiss %T on the head."), i18n.T("%N kisses %T on the head."), i18n.T("%N kisses you on the head.")},
	"nose":  {false, i18n.T("You kiss %T on the nose."), i18n.T("%N kisses %T on the nose."), i18n.T("%N kisses you on the nose.")},
	"ears":  {false, i18n.T("You kiss %T on the ears."), i18n.T("%N kisses %T on the ears."), i18n.T("%N kisses you on the ears.")},
	"neck":  {false, i18n.T("You kiss %T on the neck."), i18n.T("%N kisses %T on the neck."), i18n.T("%N kisses you on the neck.")},
	"chest": {false, i18n.T("You kiss %T on the chest."), i18n.T("%N kisses %T on the chest."), i18n.T("%N kisses you on the chest.")},
	"hand":  {false, i18n.T("You kiss %T's hand."), i18n.T("%N kisses %T's hand."), i18n.T("%N kisses your hand.")},
	"lips":  {true, i18n.T("You kiss %T on the lips."), i18n.T("%N kisses %T on the lips."), i18n.T("%N kisses you on the lips.")},
	"navel": {true, i18n.T("You kiss %T on the navel."), i18n.T("%N kisses %T on the navel."), i18n.T("%N kisses you on the navel.")},
	"leg":   {true, i18n.T("You kiss %T on the leg."), i18n.T("%N kisses %T on the leg."), i18n.T("%N kisses you on the leg.")},
	"knee":  {true, i18n.T("You kiss %T on the knee."), i18n.T("%N kisses %T on the knee."), i18n.T("%N kisses you on the knee.")},
	"feet":  {true, i18n.T("You kiss %T's feet."), i18n.T("%N kisses %T's feet."), i18n.T("%N kisses your feet.")},
}

// expandEmote replaces %N, %P, %O, %E, %T placeholders in emote strings.
func expandEmote(template string, actor *Player, targetName string) string {
	result := template
	for i := 0; i < len(result); i++ {
		if result[i] == '%' && i+1 < len(result) {
			var replacement string
			switch result[i+1] {
			case 'N':
				replacement = actor.FirstName
			case 'P':
				replacement = actor.Possessive()
			case 'O':
				replacement = actor.Objective()
			case 'E':
				replacement = actor.Pronoun()
			case 'T':
				replacement = targetName
			default:
				continue
			}
			result = result[:i] + replacement + result[i+2:]
			i += len(replacement) - 1
		}
	}
	return result
}

// processEmote handles emote commands using the emote table.
func (e *GameEngine) processEmote(player *Player, verb string, args []string) *CommandResult {
	// Emotes always reveal you
	if player.Hidden {
		player.Hidden = false
	}

	// Check race-specific emotes first
	raceKey := fmt.Sprintf("%d:%s", player.Race, verb)
	if raceEntry, ok := raceEmotes[raceKey]; ok {
		selfMsg := expandEmote(raceEntry.Self, player, "")
		roomMsg := expandEmote(raceEntry.Room, player, "")
		return &CommandResult{Messages: []string{selfMsg}, RoomBroadcast: []string{roomMsg}}
	}

	// Check additional self-emotes (no-arg only)
	if len(args) == 0 {
		if selfEmote, ok := additionalSelfEmotes[verb]; ok {
			selfMsg := expandEmote(selfEmote[0], player, "")
			roomMsg := expandEmote(selfEmote[1], player, "")
			return &CommandResult{Messages: []string{selfMsg}, RoomBroadcast: []string{roomMsg}}
		}
	}

	entry, ok := emoteTable[verb]
	if !ok {
		// Fallback generic
		v := strings.ToLower(verb)
		return &CommandResult{
			Messages:      []string{fmt.Sprintf(i18n.T("You %s."), v)},
			RoomBroadcast: []string{fmt.Sprintf(i18n.T("%s %ss."), player.FirstName, v)},
		}
	}

	if len(args) > 0 {
		targetName := strings.ToLower(strings.Join(args, " "))

		// Check for "me"/"myself" with special self-override
		if targetName == "me" || targetName == "myself" || targetName == "self" {
			key := verb + ":me"
			if override, ok := selfOverrides[key]; ok {
				selfMsg := expandEmote(override[0], player, player.FirstName)
				roomMsg := expandEmote(override[1], player, player.FirstName)
				return &CommandResult{Messages: []string{selfMsg}, RoomBroadcast: []string{roomMsg}}
			}
			selfMsg := expandEmote(entry.Self, player, player.FirstName)
			roomMsg := expandEmote(entry.Room, player, player.FirstName)
			return &CommandResult{Messages: []string{selfMsg}, RoomBroadcast: []string{roomMsg}}
		}

		// KISS with body part qualifier: "kiss <person> <bodypart>"
		if verb == "KISS" {
			return e.processKiss(player, args)
		}

		// LICK with submit-dependent behavior
		if verb == "LICK" {
			return e.processLick(player, args)
		}

		// If targeted emote templates exist
		if entry.SelfTarget != "" && entry.RoomTarget != "" {
			// Try to resolve as a player in the room
			found := e.findPlayerInRoom(player, targetName)
			if found != nil {
				displayTarget := found.FirstName
				selfMsg := expandEmote(entry.SelfTarget, player, displayTarget)
				roomMsg := expandEmote(entry.RoomTarget, player, displayTarget)
				targetMsg := expandEmote(entry.RoomTarget, player, "you")
				return &CommandResult{
					Messages:      []string{selfMsg},
					RoomBroadcast: []string{roomMsg},
					TargetName:    found.FirstName,
					TargetMsg:     []string{targetMsg},
				}
			}

			// Try to resolve as a monster in the room (including dead)
			if _, monDef := e.findMonsterInRoomIncludeDead(player, targetName); monDef != nil {
				monName := FormatMonsterName(monDef, e.monAdjs)
				displayTarget := articleFor(monName, monDef.Unique) + monName
				selfMsg := expandEmote(entry.SelfTarget, player, displayTarget)
				roomMsg := expandEmote(entry.RoomTarget, player, displayTarget)
				return &CommandResult{Messages: []string{selfMsg}, RoomBroadcast: []string{roomMsg}}
			}

			// Try to resolve as a room item
			room := e.rooms[player.RoomNumber]
			if room != nil {
				for _, ri := range room.Items {
					itemDef := e.items[ri.Archetype]
					if itemDef == nil {
						continue
					}
					name := e.getItemNounName(itemDef)
					if matchesTarget(name, targetName, e.getAdjName(ri.Adj1)) {
						displayTarget := e.formatItemName(itemDef, ri.Adj1, ri.Adj2, ri.Adj3)
						selfMsg := expandEmote(entry.SelfTarget, player, displayTarget)
						roomMsg := expandEmote(entry.RoomTarget, player, displayTarget)
						return &CommandResult{Messages: []string{selfMsg}, RoomBroadcast: []string{roomMsg}}
					}
				}
			}

			// Also check player's items (inventory + worn + wielded)
			allItems := make([]InventoryItem, 0)
			allItems = append(allItems, player.Inventory...)
			allItems = append(allItems, player.Worn...)
			if player.Wielded != nil {
				allItems = append(allItems, *player.Wielded)
			}
			for _, ii := range allItems {
				itemDef := e.items[ii.Archetype]
				if itemDef == nil {
					continue
				}
				name := e.getItemNounName(itemDef)
				if matchesTarget(name, targetName, e.getAdjName(ii.Adj1)) || matchesTarget(name, targetName, e.getAdjName(ii.Adj3)) {
					displayTarget := e.formatItemName(itemDef, ii.Adj1, ii.Adj2, ii.Adj3)
					selfMsg := expandEmote(entry.SelfTarget, player, displayTarget)
					roomMsg := expandEmote(entry.RoomTarget, player, displayTarget)
					return &CommandResult{Messages: []string{selfMsg}, RoomBroadcast: []string{roomMsg}}
				}
			}

			// Nothing matched
			return &CommandResult{Messages: []string{fmt.Sprintf(i18n.T("You don't see '%s' here."), targetName)}}
		}
	}

	selfMsg := expandEmote(entry.Self, player, "")
	roomMsg := expandEmote(entry.Room, player, "")
	return &CommandResult{Messages: []string{selfMsg}, RoomBroadcast: []string{roomMsg}}
}

// processKiss handles KISS with body part qualifiers.
// "kiss <person>" — basic kiss
// "kiss <person> <bodypart>" — kiss on specific body part (some require submit)
func (e *GameEngine) processKiss(player *Player, args []string) *CommandResult {
	if len(args) == 0 {
		return &CommandResult{Messages: []string{i18n.T("You blow a kiss.")}, RoomBroadcast: []string{fmt.Sprintf(i18n.T("%s blows a kiss."), player.FirstName)}}
	}

	// Check if last arg is a body part
	var targetArgs []string
	var bodyPart string
	if len(args) >= 2 {
		lastArg := strings.ToLower(args[len(args)-1])
		if _, ok := kissBodyParts[lastArg]; ok {
			bodyPart = lastArg
			targetArgs = args[:len(args)-1]
		}
	}
	if bodyPart == "" {
		targetArgs = args
	}

	targetName := strings.ToLower(strings.Join(targetArgs, " "))
	found := e.findPlayerInRoom(player, targetName)
	if found == nil {
		// Fall back to basic emote behavior for non-players
		entry := emoteTable["KISS"]
		if entry.SelfTarget != "" {
			// Try room items, inventory, etc.
			room := e.rooms[player.RoomNumber]
			if room != nil {
				for _, ri := range room.Items {
					itemDef := e.items[ri.Archetype]
					if itemDef == nil {
						continue
					}
					name := e.getItemNounName(itemDef)
					if matchesTarget(name, targetName, e.getAdjName(ri.Adj1)) {
						displayTarget := e.formatItemName(itemDef, ri.Adj1, ri.Adj2, ri.Adj3)
						selfMsg := expandEmote(entry.SelfTarget, player, displayTarget)
						roomMsg := expandEmote(entry.RoomTarget, player, displayTarget)
						return &CommandResult{Messages: []string{selfMsg}, RoomBroadcast: []string{roomMsg}}
					}
				}
			}
		}
		return &CommandResult{Messages: []string{fmt.Sprintf(i18n.T("You don't see '%s' here."), targetName)}}
	}

	displayTarget := found.FirstName

	if bodyPart != "" {
		bp := kissBodyParts[bodyPart]
		if bp.requiresSubmit && !found.Submitting {
			return &CommandResult{Messages: []string{fmt.Sprintf(i18n.T("%s is not allowing that."), displayTarget)}}
		}
		selfMsg := expandEmote(bp.selfMsg, player, displayTarget)
		roomMsg := expandEmote(bp.roomMsg, player, displayTarget)
		targetMsg := expandEmote(bp.targetMsg, player, displayTarget)
		return &CommandResult{
			Messages:      []string{selfMsg},
			RoomBroadcast: []string{roomMsg},
			TargetName:    found.FirstName,
			TargetMsg:     []string{targetMsg},
		}
	}

	// Basic kiss (no body part)
	selfMsg := fmt.Sprintf(i18n.T("You kiss %s."), displayTarget)
	roomMsg := fmt.Sprintf(i18n.T("%s kisses %s."), player.FirstName, displayTarget)
	targetMsg := fmt.Sprintf(i18n.T("%s kisses you."), player.FirstName)
	return &CommandResult{
		Messages:      []string{selfMsg},
		RoomBroadcast: []string{roomMsg},
		TargetName:    found.FirstName,
		TargetMsg:     []string{targetMsg},
	}
}

// processLick handles LICK with submit-dependent behavior.
// lick <person> — "You kiss <person> passionately." (if not submitted)
// lick <person> — "You lick <person> all over their body." (if submitted)
func (e *GameEngine) processLick(player *Player, args []string) *CommandResult {
	if len(args) == 0 {
		return &CommandResult{
			Messages:      []string{i18n.T("You lick your lips.")},
			RoomBroadcast: []string{fmt.Sprintf(i18n.T("%s licks %s lips."), player.FirstName, player.Possessive())},
		}
	}

	targetName := strings.ToLower(strings.Join(args, " "))
	found := e.findPlayerInRoom(player, targetName)
	if found != nil {
		displayTarget := found.FirstName
		if found.Submitting {
			selfMsg := fmt.Sprintf(i18n.T("You lick %s all over their body."), displayTarget)
			roomMsg := fmt.Sprintf(i18n.T("%s licks %s all over their body."), player.FirstName, displayTarget)
			targetMsg := fmt.Sprintf(i18n.T("%s licks you all over your body."), player.FirstName)
			return &CommandResult{
				Messages:      []string{selfMsg},
				RoomBroadcast: []string{roomMsg},
				TargetName:    found.FirstName,
				TargetMsg:     []string{targetMsg},
			}
		}
		selfMsg := fmt.Sprintf(i18n.T("You kiss %s passionately."), displayTarget)
		roomMsg := fmt.Sprintf(i18n.T("%s kisses %s passionately."), player.FirstName, displayTarget)
		targetMsg := fmt.Sprintf(i18n.T("%s kisses you passionately."), player.FirstName)
		return &CommandResult{
			Messages:      []string{selfMsg},
			RoomBroadcast: []string{roomMsg},
			TargetName:    found.FirstName,
			TargetMsg:     []string{targetMsg},
		}
	}

	// Fall back to standard emote item targeting
	entry := emoteTable["LICK"]
	room := e.rooms[player.RoomNumber]
	if room != nil && entry.SelfTarget != "" {
		for _, ri := range room.Items {
			itemDef := e.items[ri.Archetype]
			if itemDef == nil {
				continue
			}
			name := e.getItemNounName(itemDef)
			if matchesTarget(name, targetName, e.getAdjName(ri.Adj1)) {
				displayTarget := e.formatItemName(itemDef, ri.Adj1, ri.Adj2, ri.Adj3)
				selfMsg := expandEmote(entry.SelfTarget, player, displayTarget)
				roomMsg := expandEmote(entry.RoomTarget, player, displayTarget)
				return &CommandResult{Messages: []string{selfMsg}, RoomBroadcast: []string{roomMsg}}
			}
		}
	}
	return &CommandResult{Messages: []string{fmt.Sprintf(i18n.T("You don't see '%s' here."), targetName)}}
}
