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
	visitorsThisMonth: number;
	content: { [key: string]: string };
}

export type { User, Website };
