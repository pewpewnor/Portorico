"use client";

import Error from "@/components/layouts/Error";
import Loading from "@/components/layouts/Loading";
import client from "@/lib/axios";
import { getTemplateForEditing } from "@/templates/templates";
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
				console.log(data.content);

				if (res.status === 400) {
					console.log("error 400");
				} else if (res.status === 200) {
					// setWebsite({ ...data, content: JSON.parse(data.content) });
					setWebsite(data);
				}
			} catch (error) {
				setIsForbidden(true);
			}
		})();
	}, [name]);

	async function changeField(fieldName: string, value: string) {
		console.log(fieldName, value);
		if (website !== null) {
			setWebsite((prev) => {
				if (prev === null) {
					return prev;
				}
				return {
					...prev,
					content: { ...prev.content, [fieldName]: value },
				};
			});
			try {
				const jsonContent = JSON.stringify({
					...website.content,
					[fieldName]: value,
				});
				console.log(jsonContent);
				const res = await client.patch("/authed/website", {
					content: jsonContent,
					websiteId: website.id,
				});

				if (res.status === 400) {
					console.error("error 400");
				}
			} catch (error) {
				setIsForbidden(true);
			}
		}
	}

	if (isForbidden) {
		return (
			<Error
				topMessage="You need to be logged in to visit this page"
				bottomMessage="Sorry about that, please visit our login page to sign in."
				buttonText="Take me there!"
			/>
		);
	}
	if (website === null) return <Loading />;

	return getTemplateForEditing(
		website.templateName,
		website.content,
		changeField
	);
}
