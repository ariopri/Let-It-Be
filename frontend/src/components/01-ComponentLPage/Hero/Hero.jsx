import {
  Box,
  Button,
  Container,
  Flex,
  Heading,
  Image,
  Stack,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import { useCallback, useState } from 'react';
import { Link as LinkTo } from 'react-router-dom';
import { ENDPOINT_API_POST_USER } from '../../../api/api';
import useLoginState from '../../../zustand/todoLogin';
import axios from 'axios';
import { useEffect } from 'react';

export default function Hero() {
  const { isLoggedIn, dataId } = useLoginState();
  const [, setUser] = useState({});
  const color = useColorModeValue('white', 'black');
  const color2 = useColorModeValue('black', 'white');
  const bacg = useColorModeValue('accentLight.400', 'accentDark.400');
  const hoverBg = useColorModeValue('accentLight.500', 'accentDark.500');

  const getUser = useCallback(async () => {
    const headers = {
      Authorization: 'Bearer ' + localStorage.getItem('tokenId'),
    };
    const res = await axios.get(ENDPOINT_API_POST_USER + `${dataId}`, {
      headers,
    });
    setUser(res.data);
  }, [dataId]);

  useEffect(() => {
    getUser();
  }, [getUser]);
  return (
    <Stack
      as={Container}
      direction={{ base: 'column', lg: 'row' }}
      maxW={'7xl'}
    >
      <Flex
        py={{ base: 0, md: 8 }}
        flex={1}
        align={'center'}
        order={{ base: 2, lg: 1 }}
        data-aos="fade-right"
      >
        {isLoggedIn ? (
          <Stack spacing={6} w={'full'}>
            <Heading
              mt={{ base: 4, md: 0 }}
              fontSize={{ base: 'xl', md: '2xl', lg: '3xl' }}
              fontWeight={600}
            >
              <Text fontSize={{ base: '2xl', md: '3xl', lg: '3xl' }}>
                Selamat Datang,{' '}
                {`${
                  dataId.nama_depan.charAt(0).toUpperCase() +
                  dataId.nama_depan.slice(1)
                } ${
                  dataId.nama_belakang.charAt(0).toUpperCase() +
                  dataId.nama_belakang.slice(1)
                }`}
              </Text>
              <Text fontSize={'md'} color={color2} align="justify" pt={6}>
                Let It Be Akan Membantumu Mempersiapkan Diri Untuk Masuk
                Perguruan Tinggi Impian Mu.
              </Text>
            </Heading>
            <Stack
              align={'center'}
              direction={{ base: 'column-reverse', sm: 'row', md: 'row' }}
              justifyContent={{
                base: 'center',
                sm: 'space-between',
                md: 'space-between',
              }}
              spacing={{ sm: 20 }}
            >
              <Box></Box>
            </Stack>
          </Stack>
        ) : (
          <Stack spacing={6} w={'full'}>
            <Heading
              mt={{ base: 4, md: 0 }}
              fontSize={{ base: 'xl', md: '2xl', lg: '3xl' }}
              fontWeight={600}
            >
              <Text
                fontWeight="extrabold"
                fontSize={{ base: '3xl', md: '4xl', lg: '5xl' }}
              >
                Let It Be
              </Text>
              <Text fontSize={{ base: '2xl', md: '3xl', lg: '3xl' }}>
                Solusi untuk persiapan sbmptn dan Raih Mimpimu
              </Text>
            </Heading>
            <Text fontSize={'md'} color={'gray.500'} align="justify">
              Apakah Kamu Sudah Siap Untuk Mencapai Mimpimu?
              <Text fontSize={'md'} color={'gray.500'} align="justify">
                Let It Be Akan Membantumu Mempersiapkan Diri Untuk Tes Masuk
                Perguruan Tinggi. Ayo Daftarkan Dirimu Sekarang!
              </Text>
            </Text>

            <Stack
              align={'center'}
              direction={{ base: 'column-reverse', sm: 'row', md: 'row' }}
              justifyContent={{
                base: 'center',
                sm: 'space-between',
                md: 'space-between',
              }}
              spacing={{ sm: 20 }}
            >
              <Box>
                <Button
                  as={LinkTo}
                  to="/mendaftar"
                  size={{ base: 'md', md: 'lg' }}
                  rounded={'md'}
                  w={{ base: 'xs', sm: 'full' }}
                  color={color}
                  bg={bacg}
                  _hover={{
                    bg: { hoverBg },
                    transform: 'translateY(2px)',
                    boxShadow: 'lg',
                  }}
                >
                  Daftar Sekarang
                </Button>
              </Box>
            </Stack>
          </Stack>
        )}
      </Flex>
      <Flex
        data-aos="fade-left"
        flex={1}
        align={'center'}
        justify={'center'}
        order={{ base: 1, lg: 2 }}
        pl={{ md: '40px' }}
      >
        <Image
          padding={5}
          alt={''}
          objectFit={'cover'}
          src={`${process.env.PUBLIC_URL + `/image/Hero.png`}`}
        />
      </Flex>
    </Stack>
  );
}
