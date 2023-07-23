package _const

// id[ 1 к N]
var CommonSectorChanceResources = map[int]int{
	4: 3,     // thorium_ore 	100% 	100  price: 10
	1: 3,     // copper_ore  	80%  	150  price: 15
	2: 3,     // iron_ore	 	60%  	200  price: 20
	3: 75,    // silicon_ore 	40%  	250  price: 40
	5: 75,    // titanium_ore 	20% 	300  price: 50
	6: 99999, // oil			50%		100  price: 50
}

// id[1 к N]
var FreeSectorChanceResources = map[int]int{
	4: 25, // thorium_ore
	1: 25, // copper_ore
	2: 25, // iron_ore
	3: 3,  // silicon_ore
	5: 3,  // titanium_ore
	6: 3,  // oil
}

// id[1 к N]
var BattleSectorChanceResources = map[int]int{
	4: 25,  // thorium_ore
	1: 25,  // copper_ore
	2: 25,  // iron_ore
	3: 3,   // silicon_ore
	5: 3,   // titanium_ore
	6: 150, // oil
}
