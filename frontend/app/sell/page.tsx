"use client"
import { title } from "@/components/primitives";
import { Table, TableHeader, TableBody, TableColumn, TableRow, TableCell, Button } from "@nextui-org/react";
import { Spacer } from "@nextui-org/react";
import {PlusIcon} from "./PlusIcon";

export default function SellPage() {
	return (
		<>
			<div>
				<h1 className={title()} style={{ marginBottom: '20px' }}>Your listings</h1>
				<Spacer />
				<Spacer />
				<Spacer />
				<Button
					className="bg-foreground text-background"
					endContent={<PlusIcon width={undefined} height={undefined} />}
					size="sm"
				>
					Add New
				</Button>
				<Spacer />
				<Table aria-label="Example static collection table">
					<TableHeader>
						<TableColumn>NAME</TableColumn>
						<TableColumn>PRICE</TableColumn>
						<TableColumn>STATUS</TableColumn>
					</TableHeader>
					<TableBody>
						<TableRow key="1">
							<TableCell>Apple</TableCell>
							<TableCell>20</TableCell>
							<TableCell>Active</TableCell>
						</TableRow>
						<TableRow key="2">
							<TableCell>Orange</TableCell>
							<TableCell>15</TableCell>
							<TableCell>Paused</TableCell>
						</TableRow>
						<TableRow key="3">
							<TableCell>Lemon</TableCell>
							<TableCell>10</TableCell>
							<TableCell>Active</TableCell>
						</TableRow>
						<TableRow key="4">
							<TableCell>Watermelon</TableCell>
							<TableCell>20</TableCell>
							<TableCell>Active</TableCell>
						</TableRow>
					</TableBody>
				</Table>
			</div>
		</>);
}
