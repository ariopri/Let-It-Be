import {
  Button,
  Flex,
  Heading,
  SimpleGrid,
  Stack,
  Text,
  Container,
  useColorModeValue,
  Card,
  CardBody,
  CardFooter,
  Image,
  Divider,
  ButtonGroup,
} from '@chakra-ui/react';
import React from 'react';

export default function Event() {
  const color = useColorModeValue('black', 'white');
  return (
    <Stack
      as={Container}
      maxW={'7xl'}
      spacing={10}
      py={10}
      data-aos="fade-up"
      mt={10}
    >
      <Stack maxW="lg" textAlign="center" alignSelf="center">
        <Heading fontSize={{ base: 'xl', md: '2xl', lg: '3xl' }}>
          Event Kegiatan
        </Heading>
        <Text color={'gray.500'}>
          Berikut ini adalah Event Kegiatan Try Out yang akan diadakan oleh kami
          maupun oleh pihak lain
        </Text>
      </Stack>
      <Flex w="full" justifyContent={'center'} alignItems="center">
        <SimpleGrid
          w={'full'}
          columns={{ base: 2, xl: 4 }}
          spacing={4}
          color={color}
        >
          <Card maxW="sm">
            <CardBody>
              <Image
                src="https://masuk-ptn.com/images/product/5ff4dbfce4e92e2ec82d38cc9a14853963dd657b.png"
                alt=""
                borderRadius="lg"
              />
              <Stack mt="6" spacing="3">
                <Heading size="md">Try Out Utbk 2022</Heading>
                <Text>
                  Ayo ikuti try out utbk 2022 untuk mengetahui kemampuan kamu
                  dalam menghadapi UTBK 2023 nanti dan dapatkan hadiah menarik
                  dari kami
                </Text>
                <Text color="blue.600" fontSize="2xl">
                  Rp. 50.000
                </Text>
              </Stack>
            </CardBody>
            <Divider />
            <CardFooter>
              <ButtonGroup spacing="2">
                <Button variant="solid" colorScheme="blue">
                  Beli Sekarang
                </Button>
                <Button variant="ghost" colorScheme="blue">
                  Ntr Aja Deh
                </Button>
              </ButtonGroup>
            </CardFooter>
          </Card>
        </SimpleGrid>
      </Flex>
    </Stack>
  );
}
