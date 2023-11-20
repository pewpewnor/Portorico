"use client";

import { User } from "@/types/model";
import { NextUIProvider } from "@nextui-org/react";
import { ReactNode, useState } from "react";
import AuthContext from "./AuthContext";

interface ContextProviderProps {
	children: ReactNode;
}

export default function ContextProviders(props: ContextProviderProps) {
	const [user, setUser] = useState<User | null>(null);

	return (
		<NextUIProvider>
			<AuthContext.Provider value={[user, setUser]}>
				{props.children}
			</AuthContext.Provider>
		</NextUIProvider>
	);
}
