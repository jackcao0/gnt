local t = {
    ["1:1"] = { 
	    exp=100
    },
    ["1:2"] = { 
	    exp=200
    },
    ["1:3"] = { 
	    exp=300
    },
    ["2:1"] = { 
	    exp=110
    },
    ["2:2"] = { 
	    exp=210
    },
    ["2:3"] = { 
	    exp=310
    },
    ["3:1"] = { 
	    exp=150
    },
    ["3:2"] = { 
	    exp=250
    },
    ["3:3"] = { 
	    exp=400
    }
}

require'metadata'.new((...), t)
