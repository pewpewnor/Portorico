"use client";

import LoadingContext from "@/contexts/LoadingContext";
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
import { usePathname, useRouter } from "next/navigation";
import { useContext, useEffect, useState } from "react";

const menus = [
	{
		name: "Home",
		href: "/",
		hiddenIfLoggedIn: true,
	},
	{
		name: "Web Editor",
		href: "/dashboard",
	},
	{
		name: "Templates",
		href: "/templates",
	},
];

export default function Navbar() {
	const router = useRouter();
	const pathName = usePathname();

	const [isMenuOpen, setIsMenuOpen] = useState(false);
	const [isLoggedIn, setIsLoggedIn] = useState(false);
	const [isLoading, setIsLoading] = useContext(LoadingContext);

	useEffect(() => {
		setIsLoggedIn(!!Cookies.get("session"));
	}, [isLoading]);

	async function handleLogout() {
		setIsLoading(true);
		Cookies.remove("session");
		setIsLoggedIn(false);
		setIsLoading(false);
		router.push("/");
	}

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
					<NavbarItem
						key={index}
						isActive={pathName === menu.href}
						hidden={isLoggedIn && menu.hiddenIfLoggedIn}
					>
						<Link color="foreground" href={menu.href}>
							{menu.name}
						</Link>
					</NavbarItem>
				))}
			</NavbarContent>
			<NavbarContent justify="end">
				{isLoggedIn ? (
					<>
						<NavbarItem>
							<Link href="/create">
								<Button
									color="secondary"
									variant="flat"
									isDisabled={isLoading}
									className="font-bold"
								>
									Build New Website
								</Button>
							</Link>
						</NavbarItem>
						<button
							className="flex cursor-pointer"
							disabled={isLoading}
							onClick={handleLogout}
						>
							<p>Logout</p>
						</button>
					</>
				) : (
					<>
						<NavbarItem className="flex">
							<Link href="/login">Login</Link>
						</NavbarItem>
						<NavbarItem>
							<Link href="/register">
								<Button
									isDisabled={isLoading}
									color="primary"
									variant="flat"
								>
									Sign Up
								</Button>
							</Link>
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
