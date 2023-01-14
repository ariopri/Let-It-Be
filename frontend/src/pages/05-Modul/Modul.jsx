import { StarIcon } from '@chakra-ui/icons';
import {
  Button,
  Card,
  CardBody,
  CardFooter,
  CardHeader,
  Heading,
  Box,
  Image,
  SimpleGrid,
  Stack,
  Text,
  Container,
  Flex,
  Avatar,
  Link,
} from '@chakra-ui/react';

const subjects = [
  {
    name: 'Math',
    image:
      'https://storage.googleapis.com/kelase-files/kelas/66febcac2adea8f98db0064dd2f8a592.png',
    description: 'Learn about the natural world and how it works.',
    teacher: 'Mr. Smith',
    rating: 4,
    color: 'linear(to-b, rgba(246, 92, 139, 1), rgba(246, 92, 139, 1))',
  },
  {
    name: 'Math',
    image:
      'https://storage.googleapis.com/kelase-files/kelas/66febcac2adea8f98db0064dd2f8a592.png',
    description: 'Study numbers, quantities, and shapes.',
    rating: 3,
    teacher: 'Mr. John',
    color: 'linear(to-b, rgba(158, 138, 252, 1), rgba(158, 138, 252, 1))',
  },
  {
    name: 'Math',
    image:
      'https://storage.googleapis.com/kelase-files/kelas/66febcac2adea8f98db0064dd2f8a592.png',
    description: 'Explore the properties and behavior of matter.',
    rating: 5,
    teacher: 'Mr. Alex',
    color: 'linear(to-b, rgba(97, 210, 242, 1), rgba(97, 210, 242, 1))',
  },
];

export default function Modul() {
  return (
    <Link href={`/modul/${subjects.name}`}>
      <Stack
        onClick={() => {}}
        as={Container}
        maxW={'7xl'}
        spacing={10}
        py={10}
        data-aos="fade-up"
        mt={10}
      >
        <Stack
          maxW="lg"
          textAlign="center"
          alignSelf="center"
          data-aos="fade-up"
        >
          <Heading fontSize={{ base: 'xl', md: '2xl', lg: '3xl' }}>
            Daftar Kelas
          </Heading>
          <Text color={'gray.500'}>
            Beriku Beberapa Mata Pelajaran Yang Tersedia Di Let It Be
          </Text>
        </Stack>
        <Flex w="full" justifyContent={'center'} alignItems="center">
          <SimpleGrid w={'full'} columns={{ base: 2, xl: 5 }} spacing={4}>
            {subjects.map((subject, i) => (
              <Card key={i}>
                <Avatar
                  width={'100px'}
                  alignSelf={'center'}
                  m={{ base: '0 0 30px 0', md: '20px 20px 0 10px' }}
                  height={'100px'}
                  src={subject.image}
                />
                <CardBody>
                  <Heading as="h4" size="md">
                    {subject.name}
                    <Text>{subject.teacher}</Text>
                  </Heading>
                  <Text>{subject.description}</Text>
                </CardBody>
                <CardFooter>
                  <Stack isInline align="center">
                    {[...Array(5)].map((star, i) => {
                      const ratingValue = i + 1;
                      return (
                        <StarIcon
                          key={i}
                          color={
                            ratingValue <= subject.rating ? 'orange' : 'gray'
                          }
                        />
                      );
                    })}
                    <Text>{subject.rating}/5</Text>
                  </Stack>
                </CardFooter>
              </Card>
            ))}
          </SimpleGrid>
        </Flex>
      </Stack>
    </Link>
  );
}
