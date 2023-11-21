"use client";

import { NextUIProvider } from "@nextui-org/react";
import { ReactNode } from "react";

interface ContextProviderProps {
	children: ReactNode;
}

export default function ContextProviders(props: ContextProviderProps) {
	return <NextUIProvider>{props.children}</NextUIProvider>;
}
