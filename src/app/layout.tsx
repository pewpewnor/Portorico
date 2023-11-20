import Navbar from "@/components/layouts/Navbar";
import ContextProviders from "@/contexts/ContextProviders";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
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

export default function RootLayout(props: RootLayoutProps) {
	return (
		<html lang="en">
			<body className={inter.className}>
				<ContextProviders>
					<div className="flex min-h-screen w-full flex-col bg-gradient-to-b from-slate-50 to-slate-300">
						<Navbar />
						{props.children}
					</div>
				</ContextProviders>
			</body>
		</html>
	);
}
