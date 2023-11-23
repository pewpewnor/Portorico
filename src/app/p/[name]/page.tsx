"use client";

import Error from "@/components/layouts/Error";
import Loading from "@/components/layouts/Loading";
import client from "@/lib/axios";
import { getTemplate } from "@/templates/templates";
import { Website } from "@/types/model";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function WebsitePage() {
	const { name } = useParams();

	const [isNotFound, setIsNotFound] = useState(false);
	const [website, setWebsite] = useState<Website | null>(null);

	useEffect(() => {
		(async () => {
			try {
				const res = await client.get(
					"/website/" + encodeURIComponent(name as string)
				);
				const data = res.data as Website;

				if (res.status === 400) {
					console.log("error 400");
				} else if (res.status === 200) {
					setWebsite(data);
				}
			} catch (error) {
				setIsNotFound(true);
			}
		})();
	}, [name]);

	if (isNotFound) {
		return (
			<Error
				topMessage="Sorry about that, please visit our home page to get where you need to go."
				bottomMessage="Looks like you've found an empty space!"
				buttonText="Take me there!"
			/>
		);
	}
	if (!website) return <Loading />;

	return getTemplate(website.templateName, website.content);
}
