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
} from '@chakra-ui/react';

const subjects = [
  {
    name: 'Science',
    image:
      'https://storage.googleapis.com/kelase-files/kelas/66febcac2adea8f98db0064dd2f8a592.png',
    description: 'Learn about the natural world and how it works.',
    rating: 4,
    color: 'linear(to-b, rgba(246, 92, 139, 1), rgba(246, 92, 139, 1))',
  },
  {
    name: 'Math',
    image:
      'https://storage.googleapis.com/kelase-files/kelas/66febcac2adea8f98db0064dd2f8a592.png',
    description: 'Study numbers, quantities, and shapes.',
    rating: 3,
    color: 'linear(to-b, rgba(158, 138, 252, 1), rgba(158, 138, 252, 1))',
  },
  {
    name: 'Chemistry',
    image:
      'https://storage.googleapis.com/kelase-files/kelas/66febcac2adea8f98db0064dd2f8a592.png',
    description: 'Explore the properties and behavior of matter.',
    rating: 5,
    color: 'linear(to-b, rgba(97, 210, 242, 1), rgba(97, 210, 242, 1))',
  },
];

export default function Modul() {
  return (
    <Stack
      as={Container}
      maxW={'7xl'}
      spacing={10}
      py={10}
      data-aos="fade-up"
      mt={10}
    >
      <Stack maxW="lg" textAlign="center" alignSelf="center" data-aos="fade-up">
        <Heading fontSize={{ base: 'xl', md: '2xl', lg: '3xl' }}>
          Daftar Kelas
        </Heading>
        <Text color={'gray.500'}>
          Beriku Beberapa Mata Pelajaran Yang Tersedia Di Let It Be
        </Text>
      </Stack>
      <Flex w="full" justifyContent={'center'} alignItems="center">
        <SimpleGrid w={'full'} columns={{ base: 2, xl: 3 }} spacing={4}>
          {subjects.map((subject, i) => (
            <Stack
              key={i}
              justify={'space-between'}
              w="100%"
              rounded={'xl'}
              bgGradient={subject.color}
              px="8"
              py="8"
              data-aos="fade-up"
            >
              <Flex w="full" direction="column" alignItems="flex-start">
                <Text
                  fontSize={{ base: '18', md: '24' }}
                  fontWeight="500"
                  color={'black'}
                  textAlign={'left'}
                >
                  {subject.name}
                </Text>
                <Flex alignItems="flex-end" justifyContent="flex-end">
                  {[...Array(5)].map((star, i) => {
                    const ratingValue = i + 1;
                    return (
                      <Box key={i}>
                        <StarIcon
                          color={
                            ratingValue <= subject.rating ? 'orange' : 'gray'
                          }
                        />
                      </Box>
                    );
                  })}
                </Flex>
              </Flex>
              <Box>
                <Text fontSize={14} mt={4}>
                  {subject.description}
                </Text>
                <Button
                  mt={4}
                  variantColor={'teal'}
                  variant="outline"
                  size="sm"
                  textTransform={'uppercase'}
                >
                  Lihat Kelas
                </Button>
              </Box>
            </Stack>
          ))}
        </SimpleGrid>
      </Flex>
    </Stack>
  );
}
