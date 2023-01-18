import { Image, Container, Heading, Stack } from '@chakra-ui/react';

const NotFound = () => {
  return (
    <Stack
      as={Container}
      align={'center'}
      justify={'center'}
      maxW={'7xl'}
      textAlign={'center'}
    >
      <Image
        src={`${process.env.PUBLIC_URL + `/image/Img404.svg`}`}
        alt=""
        boxSize={{ base: 'xs', md: 'sm' }}
        objectFit="cover"
      />
      <Heading>Halaman Tidak Ditemukan</Heading>
    </Stack>
  );
};

export default NotFound;
