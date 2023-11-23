import {
	Button,
	Input,
	ModalBody,
	ModalContent,
	ModalFooter,
	ModalHeader,
	Modal as NextModal,
	useDisclosure,
} from "@nextui-org/react";
import { ChangeEvent, useEffect, useState } from "react";

interface ModalProps {
	defaultText: string;
	fieldName: string;
	changeContent: (fieldName: string, value: string) => void;
}

export default function Modal(props: ModalProps) {
	const { isOpen, onOpen, onOpenChange } = useDisclosure();

	const [value, setValue] = useState(props.defaultText);

	useEffect(() => {
		onOpen();
	}, [onOpen]);

	return (
		<NextModal
			isOpen={isOpen}
			onOpenChange={onOpenChange}
			placement="top-center"
		>
			<ModalContent>
				{(onClose) => (
					<>
						<ModalHeader className="flex flex-col gap-1">
							Change {props.fieldName}
						</ModalHeader>
						<ModalBody>
							<Input
								autoFocus
								label={props.fieldName}
								variant="bordered"
								name="name"
								value={value}
								onChange={(
									event: ChangeEvent<HTMLInputElement>
								) => {
									setValue(event.target.value);
								}}
							/>
						</ModalBody>
						<ModalFooter>
							<Button
								color="danger"
								variant="flat"
								onPress={onClose}
							>
								Cancel
							</Button>
							<Button
								color="primary"
								onPress={() => {
									props.changeContent(props.fieldName, value);
								}}
							>
								Create
							</Button>
						</ModalFooter>
					</>
				)}
			</ModalContent>
		</NextModal>
	);
}
