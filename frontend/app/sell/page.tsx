"use client"
import { title } from "@/components/primitives";
import { Table, TableHeader, TableBody, TableColumn, TableRow, TableCell } from "@nextui-org/react";
import { Spacer } from "@nextui-org/react";

export default function SellPage() {
	return (
		<>
			<div>
				<h1 className={title()} style={{ marginBottom: '20px' }}>Your listings</h1>
				<Spacer />
				<Spacer />
				<Spacer />
				<Table aria-label="Example static collection table">
					<TableHeader>
						<TableColumn>NAME</TableColumn>
						<TableColumn>ROLE</TableColumn>
						<TableColumn>STATUS</TableColumn>
					</TableHeader>
					<TableBody>
						<TableRow key="1">
							<TableCell>Tony Reichert</TableCell>
							<TableCell>CEO</TableCell>
							<TableCell>Active</TableCell>
						</TableRow>
						<TableRow key="2">
							<TableCell>Zoey Lang</TableCell>
							<TableCell>Technical Lead</TableCell>
							<TableCell>Paused</TableCell>
						</TableRow>
						<TableRow key="3">
							<TableCell>Jane Fisher</TableCell>
							<TableCell>Senior Developer</TableCell>
							<TableCell>Active</TableCell>
						</TableRow>
						<TableRow key="4">
							<TableCell>William Howard</TableCell>
							<TableCell>Community Manager</TableCell>
							<TableCell>Vacation</TableCell>
						</TableRow>
					</TableBody>
				</Table>
			</div>
		</>);
}
