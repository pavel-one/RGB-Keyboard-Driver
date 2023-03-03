export namespace keyboard {
	
	export class Key {
	    address: number[];
	    name: string;
	    red: number;
	    green: number;
	    blue: number;
	
	    static createFrom(source: any = {}) {
	        return new Key(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.address = source["address"];
	        this.name = source["name"];
	        this.red = source["red"];
	        this.green = source["green"];
	        this.blue = source["blue"];
	    }
	}

}

