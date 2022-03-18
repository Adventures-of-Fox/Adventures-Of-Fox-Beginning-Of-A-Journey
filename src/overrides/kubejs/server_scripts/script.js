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
