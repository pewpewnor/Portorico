interface Base {
	id: string;
	createdAt: Date;
	updatedAt: Date;
}

interface User extends Base {
	username: string;
	password: string;
}

interface Website extends Base {
	name: string;
	templateName: string;
	description: string;
	visitorsThisMonth: number;
	content: any;
}

export type { User, Website };
