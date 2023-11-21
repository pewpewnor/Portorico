"use client";

import { NextUIProvider } from "@nextui-org/react";
import Cookies from "js-cookie";
import { ReactNode, useEffect, useState } from "react";
import LoggedInContext from "./LoggedInContext";

interface ContextProviderProps {
	children: ReactNode;
}

export default function ContextProviders(props: ContextProviderProps) {
	const [isLoggedIn, setIsLoggedIn] = useState(false);

	function refreshLoggedIn() {
		setIsLoggedIn(!!Cookies.get("session"));
	}

	useEffect(() => {
		refreshLoggedIn();
	}, []);

	return (
		<NextUIProvider>
			<LoggedInContext.Provider value={[isLoggedIn, refreshLoggedIn]}>
				{props.children}
			</LoggedInContext.Provider>
		</NextUIProvider>
	);
}
