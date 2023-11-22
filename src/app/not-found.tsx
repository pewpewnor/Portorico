import Error from "@/components/layouts/Error";

export default function NotFoundPage() {
	return (
		<Error
			topMessage="Sorry about that, please visit our home page to get where you need to go."
			bottomMessage="Looks like you've found an empty space!"
			buttonText="Take me there!"
		/>
	);
}
