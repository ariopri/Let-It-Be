import {
  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Heading,
  SimpleGrid,
  Text,
} from '@chakra-ui/react';

const subjects = [
  {
    name: 'Science',
    image: 'science-image-url-goes-here',
    description: 'Learn about the natural world and how it works.',
  },
  {
    name: 'Math',
    image: 'math-image-url-goes-here',
    description: 'Study numbers, quantities, and shapes.',
  },
  {
    name: 'Chemistry',
    image: 'chemistry-image-url-goes-here',
    description: 'Explore the properties and behavior of matter.',
  },
];

export default function Modul() {
  return (
    <SimpleGrid
      spacing={4}
      templateColumns="repeat(auto-fill, minmax(200px, 1fr))"
    >
      {subjects.map((subject, i) => (
        <Card key={i}>
          <CardHeader>
            <Heading size="md">{subject.name}</Heading>
          </CardHeader>
          <CardBody>
            <img src={subject.image} alt={subject.name} />
            <Text>{subject.description}</Text>
          </CardBody>
          <CardFooter>
            <Button>View here</Button>
          </CardFooter>
        </Card>
      ))}
    </SimpleGrid>
  );
}
