export namespace config {
	
	export class Task {
	    urls?: string[];
	    method?: string;
	    timeLimit?: number;
	    turnLimit?: number;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.urls = source["urls"];
	        this.method = source["method"];
	        this.timeLimit = source["timeLimit"];
	        this.turnLimit = source["turnLimit"];
	    }
	}
	export class Config {
	    username?: string;
	    password?: string;
	    remember?: boolean;
	    baseUrl?: string;
	    unicodeOutput?: boolean;
	    tasks: Task[];
	    gSheetTokenFile?: string;
	    gSheetCredFile?: string;
	    gSheetId?: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.password = source["password"];
	        this.remember = source["remember"];
	        this.baseUrl = source["baseUrl"];
	        this.unicodeOutput = source["unicodeOutput"];
	        this.tasks = this.convertValues(source["tasks"], Task);
	        this.gSheetTokenFile = source["gSheetTokenFile"];
	        this.gSheetCredFile = source["gSheetCredFile"];
	        this.gSheetId = source["gSheetId"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace endpoint {
	
	export class SubforumPotionsClubResponse {
	    threadReports: potion.PotionClubReport[];
	
	    static createFrom(source: any = {}) {
	        return new SubforumPotionsClubResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.threadReports = this.convertValues(source["threadReports"], potion.PotionClubReport);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace tool {
	
	export class LoginResponse {
	    success?: boolean;
	    message?: string;
	    username?: string;
	    initials?: string;
	
	    static createFrom(source: any = {}) {
	        return new LoginResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.message = source["message"];
	        this.username = source["username"];
	        this.initials = source["initials"];
	    }
	}

}

