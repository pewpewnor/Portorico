"use client";

import Loading from "@/components/layouts/Loading";
import LoadingContext from "@/contexts/LoadingContext";
import client from "@/lib/axios";
import { Website } from "@/types/model";
import {
	Button,
	Card,
	CardBody,
	CardFooter,
	CardHeader,
	Divider,
	Link,
} from "@nextui-org/react";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { useContext, useEffect, useState } from "react";

export default function DashboardPage() {
	const router = useRouter();

	const [_, setLoading] = useContext(LoadingContext);
	const [websites, setWebsites] = useState<Website[]>([]);

	useEffect(() => {
		(async () => {
			setLoading(true);
			try {
				const res = await client.get("/websites");
				const data = res.data as Website[];
				console.log(data, res.status);

				if (res.status === 400) {
				} else if (res.status === 200) {
					setWebsites(data);
				}
			} catch (error) {
				setLoading(false);
				router.replace("/login");
				return;
			}
			setLoading(false);
		})();
	}, [router, setLoading]);

	return (
		<div className="mt-8 flex flex-col gap-6">
			<h1 className="text-center text-2xl font-bold">My Websites</h1>
			<div className="grid grid-cols-3 gap-4 px-16">
				{websites.map((website, index) => (
					<Card key={index} className="max-w-[400px]">
						<CardHeader className="flex gap-3">
							<Image
								alt="nextui logo"
								height={40}
								src="/favicon.ico"
								width={40}
							/>
							<div className="flex flex-col">
								<p className="text-md">{website.name}</p>
								<p className="text-small text-default-500">
									{website.visitorsThisMonth} visitors this
									month
								</p>
							</div>
						</CardHeader>
						<Divider />
						<CardBody>
							<p>{website.description}</p>
						</CardBody>
						<Divider />
						<CardFooter className="flex justify-between">
							<Link
								href={"/p/" + website.name}
								className="text-indigo-600"
								showAnchorIcon
								underline="always"
							>
								View this website
							</Link>
							<p className="text-small text-default-500">
								{website.templateName}
							</p>
						</CardFooter>
					</Card>
				))}
			</div>
		</div>
	);
}
