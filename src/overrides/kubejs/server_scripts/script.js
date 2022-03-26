// priority: 0

settings.logAddedRecipes = true
settings.logRemovedRecipes = true
settings.logSkippedRecipes = false
settings.logErroringRecipes = true

onEvent('entity.spawned', event => {
	if (event.entity.type == 'betteranimalsplus:boar') {
		event.cancel()
	}
});

onEvent('recipes', (event) => {
	event.replaceOutput({}, "betteranimalsplus:fried_egg", "additionaladditions:fried_egg");
	event.replaceOutput({}, "vanillatweaks:fried_egg", "additionaladditions:fried_egg");
});