"use client";

import Loading from "@/components/layouts/Loading";
import { NextUIProvider } from "@nextui-org/react";
import { ReactNode, useState } from "react";
import LoadingContext from "./LoadingContext";

interface ContextProviderProps {
	children: ReactNode;
}

export default function ContextProviders(props: ContextProviderProps) {
	const [isLoading, setIsLoading] = useState(false);

	return (
		<NextUIProvider>
			<LoadingContext.Provider value={[isLoading, setIsLoading]}>
				{isLoading && <Loading />}
				{props.children}
			</LoadingContext.Provider>
		</NextUIProvider>
	);
}
