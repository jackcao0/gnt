local t = {
    ["1001:1"] = { 
	    id=1001, 
	    star=1, 
	    name="三板斧", 
	    atk=100, 
	    def=120, 
	    mp=12
    },
    ["1001:2"] = { 
	    id=1001, 
	    star=2, 
	    name="三板斧", 
	    atk=120, 
	    def=140, 
	    mp=23
    },
    ["1001:3"] = { 
	    id=1001, 
	    star=3, 
	    name="三板斧", 
	    atk=140, 
	    def=160, 
	    mp=40
    },
    ["1002:1"] = { 
	    id=1002, 
	    star=1, 
	    name="金箍棒", 
	    atk=300, 
	    def=200, 
	    mp=300
    },
    ["1002:2"] = { 
	    id=1002, 
	    star=2, 
	    name="金箍棒", 
	    atk=320, 
	    def=220, 
	    mp=400
    },
    ["1002:3"] = { 
	    id=1002, 
	    star=3, 
	    name="金箍棒", 
	    atk=380, 
	    def=240, 
	    mp=500
    }
}

require'metadata'.new((...), t)
