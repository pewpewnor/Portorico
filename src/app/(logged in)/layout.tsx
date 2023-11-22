"use client";

import Error from "@/components/layouts/Error";
import LoggedInContext from "@/contexts/LoggedInContext";
import { useContext } from "react";

interface LoggedInLayoutProps {
	children: React.ReactNode;
}

export default function LoggedInLayout(props: LoggedInLayoutProps) {
	console.log("dashboard rendered");
	const [isLoggedIn, _] = useContext(LoggedInContext);

	return isLoggedIn ? (
		props.children
	) : (
		<Error
			topMessage="You need to be logged in to visit this page"
			bottomMessage="Sorry about that, please visit our login page to sign in."
			buttonText="Take me there!"
		/>
	);
}
