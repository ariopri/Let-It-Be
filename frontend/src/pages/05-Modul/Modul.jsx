import { SearchIcon, StarIcon } from '@chakra-ui/icons';
import {
  Button,
  Card,
  CardBody,
  CardFooter,
  Heading,
  Box,
  SimpleGrid,
  Stack,
  Text,
  Container,
  Flex,
  Avatar,
  Link,
  Input,
  useColorModeValue,
  InputGroup,
  InputRightElement,
} from '@chakra-ui/react';
import axios from 'axios';
import { useEffect, useState } from 'react';
import { ENDPOINT_API_GET_MODUL } from '../../api/api';
import Loading from '../../components/02-Reusable/LoadingEffect/LoadingFetchEffect';

export default function Modul() {
  const [subjects, setSubjects] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [currentPage, setCurrentPage] = useState(1);
  const [itemsPerPage] = useState(4);
  const [searchTerm, setSearchTerm] = useState('');
  const searchstyle = {
    focusBorderColor: useColorModeValue('accentLight.400', 'accentDark.400'),
  };
  const buttonStyle = {
    color: useColorModeValue('white', 'black'),
    bg: useColorModeValue('accentLight.400', 'accentDark.400'),
    _hover: {
      bg: useColorModeValue('accentLight.500', 'accentDark.500'),
    },
  };

  useEffect(() => {
    axios.get(ENDPOINT_API_GET_MODUL).then((response) => {
      setSubjects(response.data[0].subjects);
      setIsLoading(false);
    });
  }, []);

  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentSubjects = subjects
    .filter(
      (subject) =>
        subject.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        subject.description.toLowerCase().includes(searchTerm.toLowerCase())
    )
    .slice(indexOfFirstItem, indexOfLastItem);

  return isLoading ? (
    <Loading />
  ) : (
    <Stack
      onClick={() => {}}
      as={Container}
      maxW={'7xl'}
      spacing={10}
      py={10}
      mt={10}
      data-aos="fade-up"
    >
      <Stack maxW="lg" textAlign="center" alignSelf="center">
        <Heading fontSize={{ base: 'xl', md: '2xl', lg: '3xl' }}>
          Daftar Kelas
        </Heading>
        <Text color={'gray.500'}>
          Berikut Beberapa Mata Pelajaran Yang Tersedia Di Let It Be
        </Text>

        <InputGroup>
          <Input
            rounded="full"
            type="text"
            placeholder="Cari dengan kata kunci apapun disini..."
            onChange={(e) => setSearchTerm(e.target.value)}
            {...searchstyle}
          />
          <InputRightElement width="4.75rem">
            <Button
              rounded="full"
              type="submit"
              colorScheme="blue"
              rightIcon={<SearchIcon />}
              {...buttonStyle}
            >
              Cari
            </Button>
          </InputRightElement>
        </InputGroup>
      </Stack>
      <Flex w="full" justifyContent={'center'} alignItems="center">
        <SimpleGrid
          w={'full'}
          columns={{ base: 2, xl: 4 }}
          spacing={4}
          mb={'10px'}
          mt={'10px'}
        >
          {currentSubjects.map((subject, i) =>
            subject.classes.map((classses, index) => (
              <Card
                key={classses.id}
                data-aos="fade-up"
                w={'full'}
                boxShadow={'2xl'}
                rounded={'md'}
                overflow={'hidden'}
                _hover={{
                  transform: 'translateY(-5px)',
                  boxShadow: '2xl',
                }}
                transition={'all .3s ease'}
              >
                <Avatar
                  width={'100px'}
                  alignSelf={'center'}
                  m={{ base: '0 0 30px 0', md: '20px 20px 0 10px' }}
                  height={'100px'}
                  src={subject.image}
                />
                <CardBody textAlign={'center'}>
                  <Heading as="h4" size="md">
                    {subject.name.charAt(0).toUpperCase() +
                      subject.name.slice(1)}
                  </Heading>
                  <Text>{subject.description}</Text>
                  <Text>{classses.teacher}</Text>
                  <CardBody>
                    <Flex alignItems="center" justifyContent="center">
                      <Stack isInline align="center">
                        {[...Array(5)].map((star, i) => {
                          const ratingValue = i + 1;
                          return (
                            <StarIcon
                              key={i}
                              color={
                                ratingValue <= classses.rating
                                  ? 'orange'
                                  : 'gray'
                              }
                            />
                          );
                        })}
                        <Text>{classses.rating}</Text>
                      </Stack>
                    </Flex>
                  </CardBody>
                </CardBody>
                <CardFooter>
                  <Button
                    as={Link}
                    href={`modul/${subject.name}/${classses.id}`}
                    variantColor={'teal'}
                    size={'md'}
                    m={'auto'}
                    {...buttonStyle}
                  >
                    Lihat Kelas
                  </Button>
                </CardFooter>
              </Card>
            ))
          )}
        </SimpleGrid>
      </Flex>
      <Box
        as="nav"
        textAlign="center"
        py={4}
        display={{ base: 'block', md: 'flex' }}
        alignItems="center"
        justifyContent="center"
      >
        <Button
          onClick={() => setCurrentPage(currentPage - 1)}
          disabled={currentPage === 1}
          variantColor="teal"
          mr={3}
          {...buttonStyle}
        >
          Previous
        </Button>
        <Button
          onClick={() => setCurrentPage(currentPage + 1)}
          disabled={currentSubjects.length < itemsPerPage}
          variantColor="teal"
          {...buttonStyle}
        >
          Next
        </Button>
      </Box>
    </Stack>
  );
}
