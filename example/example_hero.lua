local t = {
    [1001] = { 
	    id=1001, 
	    name="陈咬金", 
	    atk=100, 
	    def=120, 
	    mp=12, 
	    extra = {
		    {id=1001, chance=110},
		    {id=1002, chance=210}
		}, 
	    upgrade = {
		    {1001, 11},
		    {1002, 21}
		}, 
	    skills = {2001, 2003, 3004}
    },
    [1002] = { 
	    id=1002, 
	    name="孙悟空", 
	    atk=300, 
	    def=200, 
	    mp=300, 
	    extra = {
		    {id=2001, chance=220},
		    {id=2002, chance=320}
		}, 
	    upgrade = {
		    {2001, 22},
		    {2002, 32}
		}, 
	    skills = {3005}
    },
    [1003] = { 
	    id=1003, 
	    name="沙僧", 
	    atk=260, 
	    def=500, 
	    mp=120, 
	    extra = {
		    {id=3001, chance=220},
		    {id=3002, chance=320}
		}, 
	     
	    
    }
}

require'metadata'.new((...), t)
