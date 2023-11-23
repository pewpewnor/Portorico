import Image from "next/image";

interface TestTemplateProps {
	isEditing: boolean;
	text1?: string;
	text2?: string;
}

export default function TestTemplate(props: TestTemplateProps) {
	const text1 = props.text1 ?? "lorem ipsum";
	const text2 = props.text2 ?? "lorem ipsum 2";

	return (
		<div className="min-h-screen bg-gray-100">
			{/* Navigation Bar */}
			<nav className="bg-white p-4">
				<div className="mx-auto flex max-w-4xl justify-center">
					<button className="mr-4 text-gray-600 hover:text-gray-800 focus:outline-none">
						Home
					</button>
					<button className="mr-4 text-gray-600 hover:text-gray-800 focus:outline-none">
						About Me
					</button>
				</div>
			</nav>

			{/* Hero Section */}
			<div className="relative h-80 overflow-hidden" id="hero">
				<Image
					src="/images/register.webp"
					placeholder="blur"
					blurDataURL="/images/register-blur.webp"
					alt="Hero Image"
					layout="fill"
					objectFit="cover"
					objectPosition="center"
				/>
			</div>

			{/* About Me Section */}
			<div className="mx-auto max-w-4xl px-4 py-16" id="about">
				<div className="rounded-lg bg-white p-8 shadow-md">
					<h2 className="mb-4 text-3xl font-bold">About Me</h2>
					<p className="text-gray-800">{text1}</p>
				</div>
			</div>

			{/* Add more sections as needed */}
		</div>
	);
}
