import Image from "next/image";

export default function NotFound() {
	return (
		<div className="flex flex-col-reverse items-center justify-center gap-16 px-4 py-24 md:gap-28 md:px-44 md:py-20 lg:flex-row lg:px-24 lg:py-24">
			<div className="relative w-full pb-12 lg:pb-0 xl:w-1/2 xl:pt-24">
				<div className="relative">
					<div className="absolute">
						<div className="">
							<h1 className="my-2 text-2xl font-bold text-gray-800">
								Looks like you&apos;ve found an empty space!
							</h1>
							<p className="my-2 text-gray-800">
								Sorry about that, please visit our home page to
								get where you need to go.
							</p>
							<button className="md my-2 rounded border bg-indigo-600 px-8 py-4 text-center text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-700 focus:ring-opacity-50 sm:w-full lg:w-auto">
								Take me there!
							</button>
						</div>
					</div>
					<div>
						<Image
							src="/images/404-text.png"
							alt="404"
							width={450}
							height={450}
						/>
					</div>
				</div>
			</div>
			<div>
				<Image
					src="/images/404-image.png"
					alt="cannot load this page"
					width={500}
					height={500}
				/>
			</div>
		</div>
	);
}
