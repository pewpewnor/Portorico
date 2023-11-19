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
import { useState } from "react";

export default function Navbar() {
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
		<NextNavbar onMenuOpenChange={setIsMenuOpen}>
			<NavbarContent>
				<NavbarMenuToggle
					aria-label={isMenuOpen ? "Close menu" : "Open menu"}
					className="sm:hidden"
				/>
				<NavbarBrand>
					<Image
						src="/images/logo.png"
						width={26}
						height={26}
						alt="logo"
						className="mx-0 h-7 w-7 lg:mx-2"
					></Image>
					<p className="hidden text-xl text-inherit sm:block lg:text-2xl">
						portorico
					</p>
				</NavbarBrand>
			</NavbarContent>

			<NavbarContent className="hidden gap-6 sm:flex" justify="center">
				<NavbarItem>
					<Link color="foreground" href="#">
						Home
					</Link>
				</NavbarItem>
				<NavbarItem isActive>
					<Link color="foreground" href="#">
						My Websites
					</Link>
				</NavbarItem>
				<NavbarItem>
					<Link color="foreground" href="#">
						Premium
					</Link>
				</NavbarItem>
				<NavbarItem>
					<Link color="foreground" href="#">
						Help
					</Link>
				</NavbarItem>
			</NavbarContent>
			<NavbarContent justify="end">
				<NavbarItem className="flex">
					<Link href="#">Login</Link>
				</NavbarItem>
				<NavbarItem>
					<Button as={Link} color="primary" href="#" variant="flat">
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
