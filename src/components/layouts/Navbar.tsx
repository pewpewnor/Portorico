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
import Cookies from "js-cookie";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";

const menus = [
	{
		name: "Home",
		href: "/",
	},
	{
		name: "My Websites",
		href: "/dashboard",
	},
	{
		name: "Premium",
		href: "/premium",
	},
	{
		name: "Help",
		href: "/help",
	},
];

export default function Navbar() {
	const pathName = usePathname();

	const [isMenuOpen, setIsMenuOpen] = useState(false);
	const [isLoggedIn, setIsLoggedIn] = useState(false);

	useEffect(() => {
		setIsLoggedIn(!!Cookies.get("session"));
	}, []);

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
				{menus.map((menu, index) => (
					<NavbarItem key={index} isActive={pathName === menu.href}>
						<Link color="foreground" href={menu.href}>
							{menu.name}
						</Link>
					</NavbarItem>
				))}
			</NavbarContent>
			<NavbarContent justify="end">
				{!isLoggedIn && (
					<>
						<NavbarItem className="flex">
							<Link href="/login">Login</Link>
						</NavbarItem>
						<NavbarItem>
							<Button
								color="primary"
								href="/register"
								variant="flat"
							>
								Sign Up
							</Button>
						</NavbarItem>
					</>
				)}
			</NavbarContent>
			<NavbarMenu>
				{menus.map((menu, index) => (
					<NavbarMenuItem
						key={index}
						isActive={pathName === menu.href}
					>
						<Link
							color="foreground"
							className="w-full"
							href={menu.href}
						>
							{menu.name}
						</Link>
					</NavbarMenuItem>
				))}
			</NavbarMenu>
		</NextNavbar>
	);
}
