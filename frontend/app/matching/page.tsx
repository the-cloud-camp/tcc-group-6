"use client"
import React from "react";
import { Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, getKeyValue, Button } from "@nextui-org/react";

const rows = [
	{
		key: "1",
		my_product: "Apple",
		other_product: "Banana",
		action: "ACCEPT",
	},
];

const columns = [
	{
		key: "my_product",
		label: "MY_PRODUCT",
	},
	{
		key: "other_product",
		label: "OTHER_PRODUCT",
	},
	{
		key: "action",
		label: "ACTION",
	},
];

export default function App() {
	return (
		<>
			<div>
				<h1>Received</h1>
				<Table aria-label="Example table with dynamic content">
					<TableHeader columns={columns}>
						{(column) => <TableColumn key={column.key}>{column.label}</TableColumn>}
					</TableHeader>
					<TableBody items={rows}>
						{(item) => (
							<TableRow key={item.key}>
								{(columnKey) => <TableCell> {getKeyValue(item, columnKey)} </TableCell>}
							</TableRow>
						)}
					</TableBody>
				</Table>
			</div>

			<div>
				<h1>Send</h1>
				<Table aria-label="Example table with dynamic content">
					<TableHeader columns={columns}>
						{(column) => <TableColumn key={column.key}>{column.label}</TableColumn>}
					</TableHeader>
					<TableBody items={rows}>
						{(item) => (
							<TableRow key={item.key}>
								{(columnKey) => <TableCell>{getKeyValue(item, columnKey)}</TableCell>}
							</TableRow>
						)}
					</TableBody>
				</Table>
			</div>

			<div>
				<h1>Matching History</h1>
				<Table aria-label="Example table with dynamic content">
					<TableHeader columns={columns}>
						{(column) => <TableColumn key={column.key}>{column.label}</TableColumn>}
					</TableHeader>
					<TableBody items={rows}>
						{(item) => (
							<TableRow key={item.key}>
								{(columnKey) => <TableCell>{getKeyValue(item, columnKey)}</TableCell>}
							</TableRow>
						)}
					</TableBody>
				</Table>
			</div>
		</>
	);
}
