"use client";

import {
	Button,
	NavbarBrand,
	NavbarContent,
	NavbarItem,
	NavbarMenu,
	NavbarMenuItem,
	NavbarMenuToggle,
	Navbar as NextNavbar,
} from "@nextui-org/react";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useState } from "react";

export default function Navbar() {
	const pathName = usePathname();

	const [isMenuOpen, setIsMenuOpen] = useState(false);

	const menuItems = [
		"Home",
		"My Websites",
		"Premium",
		"Analytics",
		"Log Out",
		"Help & Feedback",
	];

	return (
		<NextNavbar onMenuOpenChange={setIsMenuOpen} position="sticky">
			<NavbarContent>
				<NavbarMenuToggle
					aria-label={isMenuOpen ? "Close menu" : "Open menu"}
					className="sm:hidden"
				/>
				<NavbarBrand>
					<Link href="/" className="flex">
						<Image
							src="/images/logo.png"
							width={26}
							height={26}
							alt="logo"
							className="mx-0 h-7 w-7 select-none lg:mx-2"
						></Image>
						<p className="hidden select-none text-xl text-inherit sm:block lg:text-2xl">
							portorico
						</p>
					</Link>
				</NavbarBrand>
			</NavbarContent>

			<NavbarContent className="hidden gap-6 sm:flex" justify="center">
				<NavbarItem isActive={pathName === "/"}>
					<Link color="foreground" href="/">
						Home
					</Link>
				</NavbarItem>
				<NavbarItem isActive={pathName === "/dasboard"}>
					<Link color="foreground" href="dashboard">
						My Websites
					</Link>
				</NavbarItem>
				<NavbarItem isActive={pathName === "/premium"}>
					<Link color="foreground" href="#">
						Premium
					</Link>
				</NavbarItem>
				<NavbarItem isActive={pathName === "/help"}>
					<Link color="foreground" href="#">
						Help
					</Link>
				</NavbarItem>
			</NavbarContent>
			<NavbarContent justify="end">
				<NavbarItem className="flex">
					<Link href="/login">Login</Link>
				</NavbarItem>
				<NavbarItem>
					<Button
						as={Link}
						color="primary"
						href="/register"
						variant="flat"
					>
						Sign Up
					</Button>
				</NavbarItem>
			</NavbarContent>
			<NavbarMenu>
				{menuItems.map((item, index) => (
					<NavbarMenuItem key={`${item}-${index}`}>
						<Link
							color={
								index === 2
									? "primary"
									: index === menuItems.length - 1
									  ? "danger"
									  : "foreground"
							}
							className="w-full"
							href="#"
						>
							{item}
						</Link>
					</NavbarMenuItem>
				))}
			</NavbarMenu>
		</NextNavbar>
	);
}
