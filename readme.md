# Manage Amazon connect flows from command line
This is a simple tool to manage Amazon Connect flows from the command line. It is designed to be used in a CI/CD pipeline to manage the flow of your Amazon Connect instance.



## Aufruf

```bash
flowmanager --id 45b887ca-47e4-40d2-bb14-76c11ac07112
```

Download of all flows und modules of the connect instance.
Supdirectorys are created for flows and modules in `output/$stage`.

```bash
output
└── dev
    ├── flows
    └── modules
```


### AWS cli commands

## 1 Get flows

```bash
aws connect list-contact-flows --instance-id da6a09cd-a8d6-4643-a821-57c073e864c5
```

## 2 Get flow

```bash
aws connect describe-contact-flow --instance-id "da6a09cd-a8d6-4643-a821-57c073e864c5" --contact-flow-id "c9ba8115-bbd7-4fc9-bc5a-a7b87f1e6df0"
``
cp dist/flowmanager /usr/local/bin
