console.log("shtty kubejs is here")
onEvent('recipes', (event) => {
  event.replaceOutput({}, "betteranimalsplus:fried_egg", "additionaladditions:fried_egg");
  event.replaceOutput({}, "vanillatweaks:fried_egg", "additionaladditions:fried_egg");
});

