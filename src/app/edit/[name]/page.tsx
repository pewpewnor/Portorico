"use client";

import Error from "@/components/layouts/Error";
import Loading from "@/components/layouts/Loading";
import client from "@/lib/axios";
import { getTemplate, getTemplateForEditing } from "@/templates/templates";
import { Website } from "@/types/model";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function EditPage() {
	const { name } = useParams();

	const [isForbidden, setIsForbidden] = useState(false);
	const [website, setWebsite] = useState<Website | null>(null);

	useEffect(() => {
		(async () => {
			try {
				const res = await client.get(
					"/authed/website/" + encodeURIComponent(name as string)
				);
				const data = res.data as Website;

				if (res.status === 400) {
					console.log("error 400");
				} else if (res.status === 200) {
					setWebsite(data);
				}
			} catch (error) {
				setIsForbidden(true);
			}
		})();
	}, [name]);

	if (isForbidden) {
		return (
			<Error
				topMessage="You need to be logged in to visit this page"
				bottomMessage="Sorry about that, please visit our login page to sign in."
				buttonText="Take me there!"
			/>
		);
	}
	if (!website) return <Loading />;

	return getTemplateForEditing(website.templateName, website.content);
}
