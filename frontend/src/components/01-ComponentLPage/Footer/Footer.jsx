import { Box, Text, Flex, useColorModeValue, Image } from '@chakra-ui/react';

export default function Footer() {
  return (
    <Box py={6}>
      <Flex
        align={'center'}
        _before={{
          content: '""',
          borderBottom: '1px solid',
          borderColor: useColorModeValue('gray.200', 'gray.700'),
          flexGrow: 1,
          mr: 8,
        }}
        _after={{
          content: '""',
          borderBottom: '1px solid',
          borderColor: useColorModeValue('gray.200', 'gray.700'),
          flexGrow: 1,
          ml: 8,
        }}
      >
        <Image src={`${process.env.PUBLIC_URL + `/image/logo.svg`}`} alt="" />
      </Flex>
      <Text pt={4} fontSize={'sm'} textAlign={'center'}>
        © 2022 Let It Be. All rights reserved
      </Text>
    </Box>
  );
}
