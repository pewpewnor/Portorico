import { Button } from "@nextui-org/react";
import Image from "next/image";
import Link from "next/link";

interface ErrorProps {
	topMessage: string;
	bottomMessage: string;
	buttonText: string;
}

export default function Error(props: ErrorProps) {
	return (
		<div className="flex flex-col-reverse items-center justify-center gap-16 px-10 py-24 md:gap-28 md:px-44 md:py-20 lg:flex-row lg:px-24 lg:py-24">
			<div className="relative w-full pb-12 lg:pb-0 xl:w-1/2 xl:pt-24">
				<div className="relative">
					<div className="absolute">
						<h1 className="my-2 text-2xl font-bold text-gray-800">
							{props.topMessage}
						</h1>
						<p className="my-2 text-gray-800">
							{props.bottomMessage}
						</p>
						<Link href="/login">
							<Button className="my-2 w-full rounded border bg-indigo-600 px-8 py-4 text-center text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-700 focus:ring-opacity-50 lg:w-auto">
								{props.buttonText}
							</Button>
						</Link>
					</div>
					<div>
						<Image
							src="/images/404-text.png"
							className="select-none"
							alt="403"
							width={450}
							height={450}
						/>
					</div>
				</div>
			</div>
			<div>
				<Image
					src="/images/404-image.png"
					className="select-none"
					alt="cannot load this page"
					width={500}
					height={500}
				/>
			</div>
		</div>
	);
}
