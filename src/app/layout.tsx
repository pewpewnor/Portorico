import type { Metadata } from "next";
import { Inter } from "next/font/google";
import { FC } from "react";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
	title: "Portorico",
	description:
		"Portorico is a free-to-use online tool build your own portfolio website, company website, and more.",
};

interface RootLayoutProps {
	children: React.ReactNode;
}

const RootLayout: FC<RootLayoutProps> = (props: RootLayoutProps) => {
	return (
		<html lang="en">
			<body className={inter.className}>{props.children}</body>
		</html>
	);
};

export default RootLayout;
