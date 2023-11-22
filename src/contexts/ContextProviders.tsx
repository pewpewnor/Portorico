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

	function refreshIsLoggedIn() {
		setIsLoggedIn(!!Cookies.get("session"));
	}

	useEffect(() => {
		refreshIsLoggedIn();
	}, []);

	return (
		<NextUIProvider>
			<LoggedInContext.Provider value={[isLoggedIn, refreshIsLoggedIn]}>
				{props.children}
			</LoggedInContext.Provider>
		</NextUIProvider>
	);
}
